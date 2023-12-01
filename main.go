package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"toDoBackEnd/di"
	"toDoBackEnd/di/container"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"net/http"
	infraLog "toDoBackEnd/infra/log"
	"toDoBackEnd/proto/proto-gen/pb"
)

const (
	portGRPC = 9000
)

func main() {
	ctx := context.Background()
	logger, err := infraLog.New("DEBUG")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup logger: %s\n", err)
		os.Exit(1)
	}
	if err := container.InitializeContainer(ctx, logger); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to initialize container: %s\n", err)
		os.Exit(1)
	}

	errors := make(chan error)
	go func() {
		errors <- startGRPCServer(portGRPC)
	}()

	go func() {
		errors <- startHTTPServer()
	}()

	for err := range errors {
		_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
}

func startGRPCServer(port int) error {
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	pb.RegisterUserServiceServer(s, di.NewUserServer())

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] failed to listen: %v", err)
		os.Exit(1)
	}

	fsCli := container.FirestoreClient()
	defer fsCli.Close()

	fmt.Printf("[INFO] Starting gRPC server")
	return s.Serve(lis)
}

func startHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	corsConfig := cors.Options{
		AllowedMethods: []string{"POST", "GET"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}

	r.Use(cors.New(corsConfig).Handler)
	if err := registerServiceHandlers(ctx, r); err != nil {
		return err
	}

	fmt.Printf("[INFO] Starting HTTP server")
	return http.ListenAndServe(fmt.Sprintf("localhost:%s", "9001"), r)
}

func registerServiceHandlers(ctx context.Context, router *chi.Mux) error {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":%d", portGRPC)
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}

	protoMarshaller := runtime.ProtoMarshaller{}
	muxOptions := []runtime.ServeMuxOption{
		runtime.WithMarshalerOption("application/protobuf", &protoMarshaller),
		runtime.WithMarshalerOption("application/x-protobuf", &protoMarshaller),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
				UseProtoNames:   true,
			},
		}),
	}

	mux := runtime.NewServeMux(muxOptions...)

	registerFuncs := []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		pb.RegisterUserServiceHandler,
	}

	for _, registerFunc := range registerFuncs {
		if err := registerFunc(ctx, mux, conn); err != nil {
			return err
		}
	}

	prefix := "/v1"
	router.Mount(prefix, http.StripPrefix(prefix, mux))
	return nil
}

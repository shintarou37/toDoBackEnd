package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "toDoBackEnd/proto/proto-gen/pb"
)

func main() {
	// gRPCサーバーへの接続を確立
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// SayHelloメソッドを呼び出し
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Get(ctx, &pb.UserGetRequest{Id: "test"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("user name: ", resp.Name)
}

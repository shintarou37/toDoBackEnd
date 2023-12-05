package grpc

import (
	"context"
	"toDoBackEnd/application"
	"toDoBackEnd/domain/model/todo"
	pb "toDoBackEnd/proto/proto-gen/pb"

	"go.uber.org/zap"
)

type TodoServer struct {
	logger      *zap.Logger
	todoService application.TodoService
	pb.UnimplementedTodoServiceServer
}

func NewTodoServer(
	todoService application.TodoService,
	logger *zap.Logger,
) pb.TodoServiceServer {
	return &TodoServer{
		todoService: todoService,
		logger:      logger.Named("TodoServer"),
	}
}

func (s *TodoServer) Get(ctx context.Context, req *pb.TodoGetRequest) (*pb.TodoGetResponse, error) {
	todo, err := s.todoService.Get(ctx, req.GetTodoId())
	if err != nil {
		return nil, err
	}
	return &pb.TodoGetResponse{
		Todo: todo.ToPB(),
	}, nil
}

func (s *TodoServer) List(ctx context.Context, req *pb.TodoListRequest) (*pb.TodoListResponse, error) {
	todoList, err := s.todoService.List(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.TodoListResponse{
		Todo: todo.ToPBList(todoList),
	}, nil
}

func (s *TodoServer) Create(ctx context.Context, req *pb.TodoCreateRequest) (*pb.TodoCreateResponse, error) {
	todo, err := s.todoService.Create(ctx, req.Todo.GetTitle(), req.Todo.GetContext(), req.GetUserId(), req.Todo.GetPriority())
	if err != nil {
		return nil, err
	}
	return &pb.TodoCreateResponse{
		Todo: todo.ToPB(),
	}, nil
}

func (s *TodoServer) Update(ctx context.Context, req *pb.TodoUpdateRequest) (*pb.TodoUpdateResponse, error) {
	todo, err := s.todoService.Update(ctx, req.Todo.GetTitle(), req.Todo.GetContext(), req.GetTodoId(), req.Todo.GetPriority())
	if err != nil {
		return nil, err
	}
	return &pb.TodoUpdateResponse{
		Todo: todo.ToPB(),
	}, nil
}

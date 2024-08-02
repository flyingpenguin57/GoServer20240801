package user_rpc

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

type server struct {
	UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return &GetUserResponse{UserId: req.UserId, Name: "John Doe"}, nil
}

func StartRpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterUserServiceServer(s, &server{})

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

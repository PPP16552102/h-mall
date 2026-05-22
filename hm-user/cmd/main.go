package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/h-mall/hm-user/internal/application"
	v1 "github.com/h-mall/proto-repo/api/user/v1"
	"google.golang.org/grpc"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)  {
	start := time.Now()

	log.Printf("[gRPC] Method: %s, Request: %v", info.FullMethod, req)

	resp, err := handler(ctx, req)

	log.Printf("[gRPC] Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)

	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

	v1.RegisterUserServiceServer(s, application.NewUserService())

	log.Println("User Service (gRPC) is running on :8081")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
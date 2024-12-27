package main

import (
	"log"
	"net"

	"{{cookiecutter.project_slug}}/internal/logger"
	pb "{{cookiecutter.project_slug}}/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	pb.Unimplemented{{cookiecutter.grpc_service_name}}Server
}

func main() {
	logger.InitLogger() // 初始化日志
	logger.Logger.Info("Starting gRPC server")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Logger.Fatal("Failed to listen", zap.Error(err))
	}
	s := grpc.NewServer()
	pb.Register{{cookiecutter.grpc_service_name}}Server(s, &server{})

	logger.Logger.Info("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		logger.Logger.Fatal("Failed to serve", zap.Error(err))
	}
}

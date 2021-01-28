package main

import (
	"log"
	"net"
	pb "test/pb"
	"test/service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建一个 gRPC 服务端实例
	s := grpc.NewServer()

	// 初始化 GreeterServer
	greeterServer := service.NewGreeterServer()

	pb.RegisterGreeterServer(s, greeterServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package service

import (
	"context"
	"test/pb"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func NewGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

func (server *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello,my name is " + in.GetName()}, nil
}

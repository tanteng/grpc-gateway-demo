package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "test/pb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "sg"
)

func main() {
	// 建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	// 获取命令行中传入的第一个参数作为 name 值，否则name就使用默认值
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

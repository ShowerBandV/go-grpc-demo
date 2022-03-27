package main

import (
	"context"
	"fmt"
	pb "go-grcp-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func main() {
	params := grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 3 * time.Minute})
	s := grpc.NewServer(params)
	pb.RegisterFirstServiceServer(s, &FirstService{})
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(listener)
}

type FirstService struct {
}

func (f FirstService) HelloWorld(ctx context.Context, request *pb.FirstServiceRequest) (*pb.FirstServiceResponse, error) {
	fmt.Println(request)
	return &pb.FirstServiceResponse{
		Code:   "ok",
		Data:   "this is server. copy that",
		Status: 200,
	}, nil
}

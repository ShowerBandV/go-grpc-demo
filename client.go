package main

import (
	"context"
	"fmt"
	pb "go-grcp-demo/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewFirstServiceClient(conn)
	req := &pb.FirstServiceRequest{
		Code:   "test ok",
		Data:   "client message,do you copy?",
		Status: 200,
	}
	ctx := context.TODO()
	resp, err := client.HelloWorld(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

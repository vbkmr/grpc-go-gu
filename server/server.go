package main

import (
	"log"
	"net"

	"github.com/vbkmr/grpc-go-gu/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = "9000"

func main() {

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal("Failed to listen at port: %v, %v", PORT, err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	// Enable reflection
	reflection.Register(grpcServer)

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server over port: %v, %v", PORT, err)
	}
}

package main

import (
	"log"

	"github.com/vbkmr/grpc-go-gu/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	messageRequest := chat.MessageRequest{
		Message: "Hello from the client!",
	}

	response, err := c.SendMessage(context.Background(), &messageRequest)

	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}

	log.Printf("Response from server: %s", response.Message)

}

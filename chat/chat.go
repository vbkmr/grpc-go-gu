package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SendMessage(ctx context.Context, request *MessageRequest) (*MessageResponse, error) {
	log.Printf("Received message body from client: %s", request.Message)
	return &MessageResponse{Message: "Hello from the server!"}, nil
}

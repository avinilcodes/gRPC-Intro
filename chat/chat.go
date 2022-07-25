package chat

import (
	"context"
	"log"
)

type Server struct{}
type Message struct {
	Body string
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", in.Body)
	return &Message{Body: "Hello from the server"}, nil
}

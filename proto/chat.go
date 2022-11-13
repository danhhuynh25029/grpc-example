package proto

import (
	"context"
	"log"
)

type Server struct {
}

func NewServer() CalculatorServiceServer {
	return &Server{}
}
func (s *Server) Sayhello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
func (s *Server) mustEmbedUnimplementedCalculatorServiceServer() {}

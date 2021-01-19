package protos

import (
	"context"
	"fmt"
	"log"
	"os"
)

// Server represents the gRPC server

type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	if name, err := os.Hostname(); err != nil {
		log.Printf("Message %s received (could not determine hostname)", in.Greeting)
		return &PingMessage{Greeting: "bar"}, nil
	} else {
		log.Printf("Message %s received by host %s", in.Greeting, name)
		return &PingMessage{Greeting: fmt.Sprintf("Hello from %s.", name)}, nil
	}
}

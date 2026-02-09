// internal/shared/server/grpc.go
package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer(
	// Add interceptors for logging, tracing, auth — see go-observability
	)
	reflection.Register(s) // Enable for dev/staging
	return s
}

func ServeGRPC(s *grpc.Server, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}
	return s.Serve(lis)
}

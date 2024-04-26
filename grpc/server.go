package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen for connection: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to Serve GRPC server over port 9000: %v\n", err)
	}
}

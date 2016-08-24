package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nozaq/fibo-grpc/fiborpc"
	"github.com/nozaq/fibo-grpc/server/handlers"

	"google.golang.org/grpc"
)

const port = 8888

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	impl := &handlers.FiboImpl{}

	fiborpc.RegisterLeonardoServer(grpcServer, impl)
	grpcServer.Serve(lis)
}

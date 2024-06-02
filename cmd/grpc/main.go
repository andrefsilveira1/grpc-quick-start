package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer)
	reflection.Register(grpcServer)

	var port = ":50051"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server starting on port %s", port)
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}

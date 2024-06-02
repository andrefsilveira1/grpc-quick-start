package main

import (
	"andrefsilveira/grpc-quick-start/internal/pb"
	"andrefsilveira/grpc-quick-start/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	registerService := service.NewRegisterService()
	grpcServer := grpc.NewServer()
	pb.RegisterRegisterServiceServer(grpcServer, registerService)
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

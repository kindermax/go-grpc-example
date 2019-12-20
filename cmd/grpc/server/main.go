package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"try-grpc/internal/grpc/impl"
	"try-grpc/internal/grpc/service"
)

func main() {
	netListener := getNetListener(7000)
	grpcServer := grpc.NewServer()

	repositoryServiceImpl := impl.NewRepositoryServiceGrpcImpl()
	service.RegisterRepositoryServiceServer(grpcServer, repositoryServiceImpl)

	// start the server
	log.Printf("Listen on :%d", 7000)
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return lis
}
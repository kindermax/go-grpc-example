package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"try-grpc/internal/grpc/domain"
	"try-grpc/internal/grpc/service"
)

func main() {
	serverAddress := "localhost:7000"
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Can not connect to server %s, %s", serverAddress, e)
	}
	client := service.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repositoryModel := domain.Repository{
			Id:                   int64(i),
			Name:                 string("Grpc-Demo"),
			UserId:               1245,
			IsPrivate:            true,
		}
		if responseMessage, e := client.Add(context.Background(), &repositoryModel); e != nil {
			log.Fatalf("Was not able to insert Record %v", e)
		} else {
			log.Printf("Record inserted")
			log.Println(responseMessage)
			log.Printf("==========================")
		}
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return lis
}
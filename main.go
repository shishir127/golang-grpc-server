package main

import (
	"fmt"
	"log"
	"net"

	"github.com/shishir127/golang-grpc-server/spike"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(request *spike.HelloRequest, stream spike.Streamer_SayHelloServer) error {
	err := stream.Send(&spike.HelloReply{Message: "Hello " + request.Name})
	if err != nil {
		fmt.Println("Error in sending stream")
		fmt.Println(err)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	spike.RegisterStreamerServer(s, &server{})
	fmt.Println("Starting server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

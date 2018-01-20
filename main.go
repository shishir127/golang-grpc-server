package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/shishir127/golang-grpc-server/spike"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct{}

func (s *server) SayHello(request *spike.HelloRequest, stream spike.Streamer_SayHelloServer) error {
	fmt.Println("Starting to say hello")
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		fmt.Println("Error in retrieving stream metadata")
		stream.Send(&spike.HelloReply{Message: "Server error"})
		return nil
	}
	accessToken := md["authorization"][0]
	if "" == accessToken || accessToken != "test" {
		fmt.Println("Authorization failed")
		stream.Send(&spike.HelloReply{Message: fmt.Sprintf("Sorry %s, you are not authorized", request.Name)})
		return nil
	}
	for i := 0; i < 10; i++ {
		err := stream.Send(&spike.HelloReply{Message: "Hello " + request.Name})
		if err != nil {
			fmt.Println("Error in sending stream")
			fmt.Println(err)
			return err
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("Done saying hello")
	return nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
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

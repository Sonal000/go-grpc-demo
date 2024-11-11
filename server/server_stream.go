package main

import (
	"log"
	"time"

	pb "github.com/Sonal000/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(100 * time.Second)
	}
	return nil
}

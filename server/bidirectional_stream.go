package main

import (
	"io"
	"log"

	pb "github.com/Sonal000/go-grpc/proto"
)

func (s *helloServer) SayHelloBiDirectionalStream(stream pb.GreetService_SayHelloBiDirectionalStreamServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}
		log.Printf("Got request with : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name + "from Server",
		}
		if stream.Send(res); err != nil {
			return err
		}

	}

}

package main

import (
	"io"
	"log"

	pb "github.com/Sonal000/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got an Request with : %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}

}

package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Sonal000/go-grpc/proto"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Print("BiDirectional Streaming Started")

	stream, err := client.SayHelloBiDirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send the names : %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming : %v", err)
			}
			log.Print(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while Sending : %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")

}

package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Sonal000/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Start streaming")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send the names : %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming : %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming Finished")
}

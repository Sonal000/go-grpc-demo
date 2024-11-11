package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Sonal000/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Print("client streaming started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending: %v", err)
		}
		log.Printf("send the request with : %v", name)
		time.Sleep(10 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Print("Client stream finished")
	if err != nil {
		log.Fatalf("error while recieving : %v", err)
	}

	log.Printf("%v", res.Messages)

}

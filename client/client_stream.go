package main

import (
	"context"
	pb "gRPCDemo/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Starting to SayHelloClientStreaming RPC ...")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloClientStreaming RPC: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending data to server: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Streaming Finished")
	if err != nil {
		log.Fatalf("Error while receiving response from server: %v", err)
	}
	log.Printf("Response from SayHelloClientStreaming: %v", res.GetMessages())
}

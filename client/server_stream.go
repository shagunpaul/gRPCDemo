package main

import (
	"context"
	pb "gRPCDemo/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Starting to SayHelloServerStreaming RPC ...")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Error while calling SayHelloServerStreaming RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}
		log.Printf("Response from SayHelloServerStreaming: %v", res.GetMessage())
	}
	log.Printf("Streaming finished")
}

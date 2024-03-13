package main

import (
	"context"
	pb "gRPCDemo/proto"
	"io"
	"log"
	"time"
)

func callBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Starting to SayHelloBiDiStreaming RPC ...")
	stream, err := client.SayHelloBiDiStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloBiDiStreaming RPC: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading the stream: %v", err)
			}
			log.Printf("Response from SayHelloBiDiStreaming: %v", res.GetMessage())
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending data to server: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("BiDirectionalStreaming finished")
}

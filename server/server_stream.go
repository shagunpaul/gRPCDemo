package main

import (
	pb "gRPCDemo/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: "Hello " + name}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

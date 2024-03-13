package main

import (
	pb "gRPCDemo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDiStreaming(stream pb.GreetService_SayHelloBiDiStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got message %s", req.GetName())
		res := &pb.HelloResponse{Message: "Hello " + req.GetName()}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

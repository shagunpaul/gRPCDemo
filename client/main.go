package main

import (
	pb "gRPCDemo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{Names: []string{"John", "Wick", "Bob"}}
	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callBiDirectionalStreaming(client, names)
}

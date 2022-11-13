package main

import (
	"context"
	"log"
	"source/proto"

	"google.golang.org/grpc"
)

var opts []grpc.DialOption

func main() {
	cc, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot connect port 5000")
	}
	defer cc.Close()
	client := proto.NewCalculatorServiceClient(cc)
	response, err := client.Sayhello(context.Background(), &proto.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

package main

import (
	"log"
	"source/proto"

	"google.golang.org/grpc"
)

var opts []grpc.DialOption

func main() {
	cc, err := grpc.Dial("localhost:5000", opts...)
	if err != nil {
		log.Fatal("cannot connect port 5000")
	}
	defer cc.Close()
	client := proto.CalculatorServiceClient(cc)
	log.Fatal("client running %v", client)
}

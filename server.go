package main

import (
	"log"
	"net"

	"source/proto"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("cannot open port 5000")
	}
	s := proto.Server{}
	rpc := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(rpc, &s)

	err = rpc.Serve(lis)

	if err != nil {
		log.Fatal("cannot start server")
	}
}

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
	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatal("cannot open port 5000")
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, nil)

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("cannot start server")
	}
}

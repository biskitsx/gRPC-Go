package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)
	fmt.Println("gRPC server listening on [PORT : 5050]")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

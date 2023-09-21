package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	var s *grpc.Server

	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()
	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()

	}

	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)
	fmt.Println("gRPC server listening on [PORT : 5050]")
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println("")
	}
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"client/services"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	// connect to gRPC
	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:5050", "gRPC server host")
	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		creds = insecure.NewCredentials()
	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// service
	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	// err = calculatorService.Hello("")
	err = calculatorService.Fibonacci(6)
	// err = calculatorService.Average(6, 2, 1, 52, 124)
	// err = calculatorService.Sum(6, 2, 1, 52, 124)

	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}

	}

}

package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	// err = calculatorService.Hello("Suphakit")
	err = calculatorService.Fibonacci(6)
	if err != nil {
		log.Fatal(err)
	}

}

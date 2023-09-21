package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
}
type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return &calculatorService{
		calculatorClient,
	}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service: Hello\n")
	fmt.Printf("Request: %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := base.calculatorClient.Fibonacci(ctx, &req)

	if err != nil {
		return err
	}

	fmt.Printf("Service: Fibonacci\n")
	fmt.Printf("Request: %v\n", req.N)
	for {
		res, err := stream.Recv()
		// End of file คือจบแล้ว
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response: %v\n", res.Result)
	}
	return nil
}

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
	Average(numbers ...float64) error
	Sum(number ...int32) error
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

	fmt.Printf("Service: Hello (Unary)\n")
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

	fmt.Printf("Service: Fibonacci (Server Stream)\n")
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
func (base calculatorService) Average(numbers ...float64) error {
	stream, err := base.calculatorClient.Average(context.Background())

	fmt.Printf("Service: Average (Client Stream)\n")
	fmt.Printf("Request: ")

	if err != nil {
		return err
	}
	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}

		stream.Send(&req)
		fmt.Printf("%v ", req.Number)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("\nResponse: %v\n", res.Result)
	return nil
}

func (base calculatorService) Sum(numbers ...int32) error {
	stream, err := base.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Service: Sum (Bidirection Stream)\n")

	go func() {
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			stream.Send(&req)
			fmt.Printf("Request: %v\n", req.Number)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
			}
			fmt.Printf("Response: %v\n", res.Result)
		}
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}

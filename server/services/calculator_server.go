package services

import (
	"context"
	"fmt"
	"time"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return &calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UNARY
func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

// SERVER STREAMING
func (calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for i := uint32(0); i <= req.N; i++ {
		result := fib(i)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

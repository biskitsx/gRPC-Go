package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return &calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UNARY
func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf( // ต้องใช้ status ของ "google.golang.org/grpc/status"
			codes.InvalidArgument,
			"name is required",
		)
	}
	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

// SERVER STREAM
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

// CLIENT STREAM
func (calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}

	res := AverageResponse{
		Result: sum / count,
	}

	return stream.SendAndClose(&res)
}

// BIDIRECTION STREAM
func (calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		sum += req.Number

		res := SumResponse{
			Result: sum,
		}
		stream.Send(&res)
	}
	return nil
}

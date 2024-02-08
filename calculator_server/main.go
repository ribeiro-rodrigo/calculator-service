package main

import (
	"calculator-service/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type CalculatorService struct{}

func (*CalculatorService) Sum(context context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Println("request received")

	num1 := request.GetNum1()
	num2 := request.GetNum2()

	result := num1 + num2

	fmt.Println(fmt.Sprintf("result is %d", result))

	response := calculatorpb.SumResponse{
		Result: result,
	}

	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Erro ao criar o Listen %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	calculatorpb.RegisterCalculatorServiceServer(s, &CalculatorService{})

	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	fmt.Println("Servidor rodando")

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Error %v", err)
	}
}

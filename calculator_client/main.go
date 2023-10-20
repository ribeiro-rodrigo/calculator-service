package main

import (
	"calculator-service/calculatorpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error %v", err)
	}

	defer conn.Close()

	calculatorService := calculatorpb.NewCalculatorServiceClient(conn)

	request := &calculatorpb.SumRequest{Num1: 2, Num2: 3}
	response, err := calculatorService.Sum(context.TODO(), request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("O resultado da soma é ", response.Result)

}

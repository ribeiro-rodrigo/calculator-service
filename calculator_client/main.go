package main

import (
	"calculator-service/calculatorpb"
	"context"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	arn := "arn:aws:acm:us-east-1:073521391622:certificate/add3bf34-0b92-43bd-b8e3-6943dc29a713"
	region := "us-east-1"
	address := "calculatorserver.lab.olxbr.cloud:443"

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))

	if err != nil {
		log.Fatal(err)
	}

	client := acm.NewFromConfig(cfg)
	certificate, err := client.GetCertificate(ctx, &acm.GetCertificateInput{CertificateArn: &arn})

	if err != nil {
		log.Fatal(err)
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(*certificate.Certificate))
	creds := credentials.NewClientTLSFromCert(pool, "")

	//conn, err := grpc.Dial("calculatorserver.lab.olxbr.cloud:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, address, []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}...)

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

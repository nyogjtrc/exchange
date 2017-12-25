package main

import (
	"context"
	"log"

	pb "github.com/nyogjtrc/exchange"
	"google.golang.org/grpc"
)

var serverAddr = "localhost:50001"

func GetRate(client pb.ExchangeServiceClient, req *pb.RateRequest) {
	reply, err := client.GetRate(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewExchangeServiceClient(conn)

	req := pb.RateRequest{
		Base:   "USD",
		Target: "TWD",
	}
	GetRate(client, &req)
}

package main

import (
	"context"
	"log"

	pb "github.com/nyogjtrc/exchange"
	"google.golang.org/grpc"
)

var serverAddr = "localhost:50001"

func DoGetRate(client pb.ExchangeServiceClient, base, target string) {
	req := pb.RateRequest{
		Base:   base,
		Target: target,
	}
	GetRate(client, &req)
}

func GetRate(client pb.ExchangeServiceClient, req *pb.RateRequest) {
	reply, err := client.GetRate(context.Background(), req)
	if err != nil {
		log.Println(err, req)
		return
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

	DoGetRate(client, "USD", "TWD")
	DoGetRate(client, "USD", "CCC")
	DoGetRate(client, "USD", "CNY")
	DoGetRate(client, "USD", "JPY")
	DoGetRate(client, "TWD", "JPY")

}

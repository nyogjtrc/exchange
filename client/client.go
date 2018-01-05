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

func ListRate(client pb.ExchangeServiceClient) {
	reqs := []pb.RateRequest{
		pb.RateRequest{
			Base:   "USD",
			Target: "TWD",
		},
		pb.RateRequest{
			Base:   "USD",
			Target: "CNY",
		},
		pb.RateRequest{
			Base:   "JPY",
			Target: "TWD",
		},
	}
	stream, err := client.ListRate(context.Background())
	if err != nil {
		log.Panicln(err)
		return
	}

	for _, req := range reqs {
		if err = stream.Send(&req); err != nil {
			log.Panicln(err)
			return
		}
	}

	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Panicln(err)
		return
	}

	log.Println(recv)
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

	log.Println("")

	ListRate(client)
}

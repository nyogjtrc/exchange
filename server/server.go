package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nyogjtrc/exchange"
	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:50001"
	sourceURL  = "www.google.com"
)

// ExchangeServer to implement grpc service
type ExchangeServer struct {
	sourceURL string
}

func (s *ExchangeServer) GetRate(ctx context.Context, in *pb.RateRequest) (*pb.RateReply, error) {
	log.Println(in)
	return &pb.RateReply{
		Base:   in.Base,
		Target: in.Target,
		Rate:   1.0,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExchangeServiceServer(s, &ExchangeServer{sourceURL})
	s.Serve(lis)
}

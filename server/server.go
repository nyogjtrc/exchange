package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	pb "github.com/nyogjtrc/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	serverAddr = "localhost:50001"
	sourceURL  = "www.google.com"
)

// ExchangeServer to implement grpc service
type ExchangeServer struct {
	sourceURL string
	RateMap   RateMap
}

func (s *ExchangeServer) GetRate(ctx context.Context, in *pb.RateRequest) (*pb.RateReply, error) {
	var exrate float64

	if in.Base == in.Target {
		exrate = 1
	} else if in.Base == "USD" {
		key := in.Base + in.Target
		rateData, err := s.findCurrency(key)
		if err != nil {
			return nil, err
		}
		exrate = rateData.Exrate
	} else if in.Target == "USD" {
		key := in.Target + in.Base
		rateData, err := s.findCurrency(key)
		if err != nil {
			return nil, err
		}
		exrate = 1 / rateData.Exrate
	} else {
		key := "USD" + in.Base
		baseData, err := s.findCurrency(key)
		if err != nil {
			return nil, err
		}

		key2 := "USD" + in.Target
		targetData, err := s.findCurrency(key2)
		if err != nil {
			return nil, err
		}
		exrate = targetData.Exrate / baseData.Exrate
	}

	return &pb.RateReply{
		Base:   in.Base,
		Target: in.Target,
		Rate:   exrate,
	}, nil
}

func (s *ExchangeServer) findCurrency(key string) (*Rate, error) {
	r, ok := s.RateMap[key]
	if !ok {
		return nil, grpc.Errorf(codes.OutOfRange, "currency data not found")
	}

	return &r, nil
}

type Rate struct {
	UTC    string
	Exrate float64
}

type RateMap map[string]Rate

func FetchRateAPI() RateMap {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://tw.rter.info/capi.php", nil)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	m := make(RateMap)
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func main() {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	exServer := &ExchangeServer{
		sourceURL: sourceURL,
		RateMap:   FetchRateAPI(),
	}

	s := grpc.NewServer()
	pb.RegisterExchangeServiceServer(s, exServer)
	s.Serve(lis)
}

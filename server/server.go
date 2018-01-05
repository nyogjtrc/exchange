package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/nyogjtrc/exchange"
	"github.com/nyogjtrc/exchange/health"
	"go.uber.org/zap"
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
	exrate, err := s.findRate(in.Base, in.Target)
	if err != nil {
		return nil, err
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

func (s *ExchangeServer) findRate(base, target string) (float64, error) {
	if base == target {
		return 1, nil
	} else if base == "USD" {
		key := base + target
		rateData, err := s.findCurrency(key)
		if err != nil {
			return 0, err
		}
		return rateData.Exrate, nil
	} else if target == "USD" {
		key := target + base
		rateData, err := s.findCurrency(key)
		if err != nil {
			return 0, err
		}
		return 1 / rateData.Exrate, nil
	} else {
		key := "USD" + base
		baseData, err := s.findCurrency(key)
		if err != nil {
			return 0, err
		}

		key2 := "USD" + target
		targetData, err := s.findCurrency(key2)
		if err != nil {
			return 0, err
		}
		return targetData.Exrate / baseData.Exrate, nil
	}
}

func (s *ExchangeServer) ListRate(stream pb.ExchangeService_ListRateServer) error {
	startTime := time.Now()
	var count int32
	var plys []*pb.RateReply
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RateList{
				Count:    count,
				Rates:    plys,
				CostTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}

		rate, err := s.findRate(req.Base, req.Target)
		if err != nil {
			return err
		}
		plys = append(plys, &pb.RateReply{
			Base:   req.Base,
			Target: req.Target,
			Rate:   rate,
		})

		count++
	}
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

type HealthServer struct{}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

func (s *HealthServer) Check(ctx context.Context, req *health.Empty) (*health.HealthReply, error) {
	return &health.HealthReply{
		Status: health.HealthReply_SERVING,
	}, nil
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

	healServer := NewHealthServer()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.NewExample()),
		)),
	)
	pb.RegisterExchangeServiceServer(s, exServer)
	health.RegisterHealthServer(s, healServer)
	s.Serve(lis)
}

package main

import (
	"context"
	"fmt"
	"io"
	"testing"

	pb "github.com/nyogjtrc/exchange"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func setup() *ExchangeServer {
	m := RateMap{
		"USDTWD": Rate{
			UTC:    "2017-12-28 09:12:40",
			Exrate: 29.792,
		},
		"USDJPY": Rate{
			UTC:    "2017-12-28 09:12:40",
			Exrate: 112.648003,
		},
	}
	return &ExchangeServer{
		RateMap: m,
	}
}

func TestGetRate(t *testing.T) {
	s := setup()

	testCases := []struct {
		req *pb.RateRequest
		exp *pb.RateReply
	}{
		{
			req: &pb.RateRequest{
				Base:   "USD",
				Target: "TWD",
			},
			exp: &pb.RateReply{
				Base:   "USD",
				Target: "TWD",
				Rate:   29.792,
			},
		},
		{
			req: &pb.RateRequest{
				Base:   "JPY",
				Target: "TWD",
			},
			exp: &pb.RateReply{
				Base:   "JPY",
				Target: "TWD",
				Rate:   0.26446984595013195,
			},
		},
		{
			req: &pb.RateRequest{
				Base:   "TWD",
				Target: "TWD",
			},
			exp: &pb.RateReply{
				Base:   "TWD",
				Target: "TWD",
				Rate:   1,
			},
		},
		{
			req: &pb.RateRequest{
				Base:   "TWD",
				Target: "USD",
			},
			exp: &pb.RateReply{
				Base:   "TWD",
				Target: "USD",
				Rate:   0.03356605800214823,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("base %s target %s", tc.req.Base, tc.req.Target), func(t *testing.T) {
			rep, err := s.GetRate(context.Background(), tc.req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tc.exp, rep)
		})
	}

	t.Run("not found currency", func(t *testing.T) {
		req := &pb.RateRequest{
			Base:   "USD",
			Target: "CYN",
		}

		_, err := s.GetRate(context.Background(), req)
		assert.Error(t, err)
	})
}

type MockStream struct {
	grpc.ServerStream
	Requests []pb.RateRequest
	RateList *pb.RateList
}

func (s *MockStream) SendAndClose(list *pb.RateList) error {
	s.RateList = list
	return nil
}

func (s *MockStream) Recv() (*pb.RateRequest, error) {
	if len(s.Requests) > 0 {
		var x pb.RateRequest
		x, s.Requests = s.Requests[0], s.Requests[1:]
		return &x, nil
	}
	return nil, io.EOF
}

func TestListRate(t *testing.T) {
	s := setup()

	expect := &pb.RateList{
		Count: 2,
		Rates: []*pb.RateReply{
			&pb.RateReply{
				Base:   "USD",
				Target: "TWD",
				Rate:   29.792,
			},
			&pb.RateReply{
				Base:   "USD",
				Target: "JPY",
				Rate:   112.648003,
			},
		},
	}

	t.Run("client stram", func(t *testing.T) {
		mStream := &MockStream{
			Requests: []pb.RateRequest{
				pb.RateRequest{
					Base:   "USD",
					Target: "TWD",
				},
				pb.RateRequest{
					Base:   "USD",
					Target: "JPY",
				},
			},
		}

		if assert.NoError(t, s.ListRate(mStream)) {
			assert.Equal(t, expect.Count, mStream.RateList.Count)
			assert.Equal(t, expect.Rates, mStream.RateList.Rates)
		}
	})
}

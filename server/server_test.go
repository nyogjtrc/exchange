package main

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/nyogjtrc/exchange"
	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
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

	s := ExchangeServer{
		RateMap: m,
	}

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

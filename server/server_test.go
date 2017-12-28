package main

import (
	"context"
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
	}

	s := ExchangeServer{
		RateMap: m,
	}

	t.Run("USD TWD", func(t *testing.T) {
		req := &pb.RateRequest{
			Base:   "USD",
			Target: "TWD",
		}
		expectReply := &pb.RateReply{
			Base:   req.Base,
			Target: req.Target,
			Rate:   29.792,
		}

		rep, err := s.GetRate(context.Background(), req)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expectReply, rep)
	})

	t.Run("not found currency", func(t *testing.T) {
		req := &pb.RateRequest{
			Base:   "USD",
			Target: "CYN",
		}

		_, err := s.GetRate(context.Background(), req)
		assert.Error(t, err)
	})
}

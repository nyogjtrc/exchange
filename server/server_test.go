package main

import (
	"context"
	"testing"

	pb "github.com/nyogjtrc/exchange"
	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
	s := ExchangeServer{}

	req := &pb.RateRequest{}
	expectReply := &pb.RateReply{
		Base:   req.Base,
		Target: req.Target,
		Rate:   1.0,
	}

	rep, err := s.GetRate(context.Background(), req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectReply, rep)
}

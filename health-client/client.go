package main

import (
	"context"
	"log"
	"time"

	"github.com/nyogjtrc/exchange/health"
	"google.golang.org/grpc"
)

type HealthClient struct {
	client health.HealthClient
	conn   *grpc.ClientConn
}

func NewHealthClient(conn *grpc.ClientConn) *HealthClient {
	client := new(HealthClient)
	client.client = health.NewHealthClient(conn)
	client.conn = conn
	return client
}

func (c *HealthClient) Close() error {
	return c.conn.Close()
}

func (c *HealthClient) Check() {
	req := &health.Empty{}
	reply, err := c.client.Check(context.Background(), req)
	if err != nil {
		log.Println(err, req)
		return
	}
	log.Println(reply)
}

var serverAddr = "localhost:50001"

func main() {
	for {
		conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		client := NewHealthClient(conn)
		client.Check()
		conn.Close()

		<-time.After(time.Second)
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/danielwangai/blockchain-project/node"
	"github.com/danielwangai/blockchain-project/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	node := node.NewNode()
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port 3010")
	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()
	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "bchain-0.1",
		Height:  1,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatalf("failed to handle transaction: %v", err)
	}
}

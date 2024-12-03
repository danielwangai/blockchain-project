package main

import (
	"context"
	"github.com/danielwangai/blockchain-project/node"
	"github.com/danielwangai/blockchain-project/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	makeNode(":3000", []string{})
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"})
	time.Sleep(4 * time.Second)
	makeNode(":6000", []string{":4000"})
	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr, bootstrapNodes)

	return n
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

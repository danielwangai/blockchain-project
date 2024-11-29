package main

import (
	"context"
	"github.com/danielwangai/blockchain-project/node"
	"github.com/danielwangai/blockchain-project/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	makeNode(":3000", []string{})
	makeNode(":4000", []string{":3000"})
	select {}
	//node := node.NewNode()
	//go func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		makeTransaction()
	//	}
	//}()
	//log.Fatal(node.Start(":3000"))
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr)
	if len(bootstrapNodes) > 0 {
		if err := n.BootstrapNetwork(bootstrapNodes); err != nil {
			log.Fatal(err)
		}
	}
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

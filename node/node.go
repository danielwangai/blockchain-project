package node

import (
	"context"
	"github.com/danielwangai/blockchain-project/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"sync"
)

type Node struct {
	listenAddr string
	version    string
	logger     *zap.SugaredLogger
	peerLock   sync.RWMutex
	peers      map[proto.NodeClient]*proto.Version // using a map since it's easier to manage; O(1) access
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	//loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := loggerConfig.Build()
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "bchain-0.1",
		logger:  logger.Sugar(),
	}
}

// addPeer registers nodes/peers connected to node n
func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	n.logger.Debugw("new peer connected", "addr", n.listenAddr, "height", v.Height)

	n.peers[c] = v
}

// deletePeer removes a node/peer from peer "list"
func (n *Node) deletePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) BootstrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		// create a new node for each address
		c, err := makeNodeClient(addr)
		if err != nil {
			return err
		}

		// establish connection with node n
		v, err := c.Handshake(context.Background(), n.getVersion())
		if err != nil {
			n.logger.Error("handshake error: ", err)
			continue
		}

		// register client as a peer to this node
		n.addPeer(c, v)
	}

	return nil
}

// Start spins up a new grpc server for node n
func (n *Node) Start(listenAddr string) error {
	n.listenAddr = listenAddr
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		n.logger.Errorf("failed to start: %v\n", err)
	}

	proto.RegisterNodeServer(grpcServer, n)
	n.logger.Infow("node running on", "port", listenAddr)

	return grpcServer.Serve(ln)
}

// Handshake connects a nodes to enable communication
func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	c, err := makeNodeClient(v.Version)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	n.logger.Infow("Received transaction from", "peer", peer)
	return &proto.Ack{}, nil
}

// getVersion returns the version of the node
func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "bchain-0.1",
		Height:     0,
		ListenAddr: n.listenAddr,
	}
}

// makeNodeClient spins up a new node client
func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	c, err := grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return proto.NewNodeClient(c), nil
}

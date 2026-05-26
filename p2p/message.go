package p2p

import "net"

// RPC represents a message in the p2p network
// Each transport between two nodes in the networks
type RPC struct {
	Payload []byte
	From net.Addr
}

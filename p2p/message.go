package p2p

import "net"

// Message represents a message in the p2p network
// Each transport between two nodes in the networks
type Message struct {
	Payload []byte
	From net.Addr
}

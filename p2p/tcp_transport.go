package p2p

import (
	"net"
	"sync"
)

// TCPTransport is a TCP transport
// It implements the Transport interface
// It can be used to send and receive messages over a TCP connection
type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peer map[net.Addr]Peer
}

// NewTCPTransport creates a new TCP transport
// It takes a listen address as input and returns a TCPTransport
// The listen address is the address that the transport will listen on
func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}


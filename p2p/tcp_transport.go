package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a peer in the TCP transport
// To establish a connection, a TCPPeer needs to know the address of the peer
type TCPPeer struct {
	// address is the address of the peer
	conn net.Conn
	// outbound is a boolean that indicates whether the peer is outbound or inbound
	// if we dial and retrive a connection => outboud = true
	// if we accept a connection => outbound = false
	outbound bool
}


// NewTCPPeer creates a new TCP peer
// It takes a connection and an outbound boolean as input and returns a TCPPeer
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

// TCPTransport is a TCP transport
// It implements the Transport interface
// It can be used to send and receive messages over a TCP connection
type TCPTransport struct {
	listenAddress string
	listener net.Listener

	handshakeFunc HandshakeFunc

	mu sync.RWMutex
	peer map[net.Addr]Peer
}

func NOPHandshakeFunc(any) error { return nil}
// NewTCPTransport creates a new TCP transport
// It takes a listen address as input and returns a TCPTransport
// The listen address is the address that the transport will listen on
func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

// ListenAndAccept starts the TCP transport and listens for incoming connections
// It returns an error if there is a problem with the listening
func (t *TCPTransport) ListenAndAccept()  error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConnection(conn)
	}
} 

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incomming connection %+v\n", peer)
}

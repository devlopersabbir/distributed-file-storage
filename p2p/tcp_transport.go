package p2p

import (
	"fmt"
	"net"
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

// Close closes the TCP peer
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// Send sends a message over the TCP connection
// It takes a message as input and returns an error if there is a problem with the sending
type TCPTransportOptions struct {
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Decoder Decoder

	OnPeer func(Peer) error
}

// TCPTransport is a TCP transport
// It implements the Transport interface
// It can be used to send and receive messages over a TCP connection
type TCPTransport struct {
	// TCPTransportOptions is a struct that contains the options for the TCP transport
	TCPTransportOptions
	// listener is the listener for the TCP transport
	listener net.Listener
	// rpcChan is the channel for consuming messages
	rpcChan chan RPC
}

// NewTCPTransport creates a new TCP transport
// It takes a listen address as input and returns a TCPTransport
// The listen address is the address that the transport will listen on
func NewTCPTransport(opts TCPTransportOptions) *TCPTransport{
	return &TCPTransport{
		TCPTransportOptions: opts,
		rpcChan: make(chan RPC),
	}
}

// Consume returns a channel that can be used to consume messages
// Consume implements the transport interface, which will return read-only channel 
// for reading the incomming message received from another peer
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcChan
}


// ListenAndAccept starts the TCP transport and listens for incoming connections
// It returns an error if there is a problem with the listening
func (t *TCPTransport) ListenAndAccept()  error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)
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

		fmt.Printf("New incomming connection  %+v\n", conn)
		go t.handleConnection(conn)
	}
} 


func (t *TCPTransport) handleConnection(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("Droping peer connections: %s", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil { return }

	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil { return} // TODO: handle error
	}
	// read loop
	rpc := RPC{}
	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		rpc.From = conn.RemoteAddr()

		t.rpcChan <- rpc
	}
}

package p2p

// Peer is the interface for a peer in the network
type Peer interface {

}

// Transport is the interface for a transport in the network
// A transport is a way to send and receive messages
// A transport can be a TCP connection, a UDP connection, a WebSocket connection, etc.
type Transport interface {
	ListenAndAccept() error
}

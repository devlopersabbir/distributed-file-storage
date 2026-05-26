package p2p


type HandshakeFunc func(any) error

// NOPHandshakeFunc is a handshake function that does nothing
// It is used when there is no handshake needed
func NOPHandshakeFunc(any) error { return nil}


package main

import (
	"log"

	"github.com/devlopersabbir/storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {
	}
}

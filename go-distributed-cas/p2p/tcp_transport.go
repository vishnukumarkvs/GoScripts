package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptAndLoop()
	return nil
}

func (t *TCPTransport) startAcceptAndLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection %+v\n", conn)
}

// notes
// Mutex - Its a synchronization primitive provided by sync package to protect shared resources
// using this will make sure that only one goroutine can update resource at a time which will prevent race condition

// listener
// A listener listens for incoming connections
// It can accept a connection
// once it accepts, it can also handle the connection
// We use goroutines to handle multiple connections

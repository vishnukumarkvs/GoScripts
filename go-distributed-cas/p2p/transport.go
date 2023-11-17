package p2p

// Peer is an interface which represents a remote node
type Peer interface{}

// anythings that creates communication between nodes
// TCP, UDP, Sockets etc
type Transport interface {
	ListenAndAccept() error
}

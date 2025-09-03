package simple_chat_app

import (
    "net"
    "sync"
)

// TODO:
// - Implement a basic TCP chat server:
//   - Start: listen on a TCP port and accept connections.
//   - Track clients with unique IDs and default names (e.g., Guest1).
//   - Broadcast: send messages from one client to all others.
//   - Ensure concurrent access to client map is synchronized.

type Client struct {
	conn   net.Conn
	server *Server
	id     int
	name   string
}

type Server struct {
	clients  map[int]*Client
	mu       sync.Mutex
	nextID   int
	listener net.Listener
}

func NewServer() *Server {
    // TODO: initialize server state
    return &Server{}
}

func (s *Server) Start(port string) error {
    // TODO: start listening and accept connections
    return nil
}

func (s *Server) Stop() error {
    // TODO: stop the server/listener
    return nil
}

func (s *Server) acceptConnections() {
    // TODO: accept and register clients
}

func (c *Client) handleConnection() {
    // TODO: read messages from the client and broadcast
}

func (s *Server) Broadcast(sender *Client, message string) {
    // TODO: broadcast message to other clients
}

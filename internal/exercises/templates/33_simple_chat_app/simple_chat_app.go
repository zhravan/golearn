package simple_chat_app

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

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
	return &Server{
		clients: make(map[int]*Client),
		nextID:  1,
	}
}

func (s *Server) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	s.listener = listener
	fmt.Printf("Chat server listening on :%s\n", port)

	go s.acceptConnections()
	return nil
}

func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// Listener closed or other error
			return
		}
		s.mu.Lock()
		clientID := s.nextID
		s.nextID++
		s.mu.Unlock()

		client := &Client{conn: conn, server: s, id: clientID, name: fmt.Sprintf("Guest%d", clientID)}
		s.mu.Lock()
		s.clients[clientID] = client
		s.mu.Unlock()

		fmt.Printf("Client %s connected\n", client.name)
		go client.handleConnection()
	}
}

func (c *Client) handleConnection() {
	defer func() {
		fmt.Printf("Client %s disconnected\n", c.name)
		c.server.mu.Lock()
		delete(c.server.clients, c.id)
		c.server.mu.Unlock()
		c.conn.Close()
	}()

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("[%s]: %s\n", c.name, message)
		c.server.Broadcast(c, message)
	}
}

func (s *Server) Broadcast(sender *Client, message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, client := range s.clients {
		if client.id != sender.id {
			_, err := fmt.Fprintf(client.conn, "[%s]: %s\n", sender.name, message)
			if err != nil {
				fmt.Printf("Error broadcasting to client %s: %v\n", client.name, err)
			}
		}
	}
}


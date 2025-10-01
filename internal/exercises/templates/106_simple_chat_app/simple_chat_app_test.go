package simple_chat_app

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServerStartStop(t *testing.T) {
	server := NewServer()

	err := server.Start("8082")
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}

	// Give server a moment to start
	time.Sleep(10 * time.Millisecond)

	err = server.Stop()
	if err != nil {
		t.Fatalf("Server failed to stop: %v", err)
	}
}

func TestClientConnectionAndBroadcast(t *testing.T) {
	server := NewServer()
	err := server.Start("8083")
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}
	defer server.Stop()

	time.Sleep(10 * time.Millisecond)

	// Client 1
	conn1, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		t.Fatalf("Client 1 failed to connect: %v", err)
	}
	defer conn1.Close()

	// Client 2
	conn2, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		t.Fatalf("Client 2 failed to connect: %v", err)
	}
	defer conn2.Close()

	time.Sleep(10 * time.Millisecond) // Give clients a moment to register

	// Client 1 sends message
	message1 := "Hello from client 1"
	fmt.Fprintf(conn1, message1+"\n")

	// Client 2 should receive message from client 1
	scanner2 := bufio.NewScanner(conn2)
	if !scanner2.Scan() {
		t.Fatalf("Client 2 did not receive message")
	}
	received2 := scanner2.Text()
	expected2Prefix := "[Guest1]: " + message1 // Assuming Guest1 is client 1
	if received2 != expected2Prefix {
		t.Errorf("Client 2 received unexpected message: got %q, want prefix %q", received2, expected2Prefix)
	}

	// Client 2 sends message
	message2 := "Hello from client 2"
	fmt.Fprintf(conn2, message2+"\n")

	// Client 1 should receive message from client 2
	scanner1 := bufio.NewScanner(conn1)
	if !scanner1.Scan() {
		t.Fatalf("Client 1 did not receive message")
	}
	received1 := scanner1.Text()
	expected1Prefix := "[Guest2]: " + message2 // Assuming Guest2 is client 2
	if received1 != expected1Prefix {
		t.Errorf("Client 1 received unexpected message: got %q, want prefix %q", received1, expected1Prefix)
	}
}

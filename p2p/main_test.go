package main

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		main() // Start the server
	}()

	// Allow some time for the server to start
	// You can adjust this duration as needed
	// For a real-world scenario, you may want to use a more robust approach
	// to wait for the server to start, such as polling or retries.
	// For simplicity, we'll just sleep here.
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		t.Fatalf("Failed to send data to server: %v", err)
	}

	// Wait for the server to finish handling the connection
	wg.Wait()
}

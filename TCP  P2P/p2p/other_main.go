package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func othermain() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port 8080\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection received from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	fileName, _ := reader.ReadString('\n')

	file, err := os.Open(strings.TrimSpace(fileName))

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var SERVER = "localhost:8080"


//const FILENAME = "filename.txt"

func omain() {
	fmt.Println("Client connected to", SERVER)
	conn, _ := net.Dial("tcp", SERVER)
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", FILENAME)

	reader := bufio.NewReader(conn)
	content := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		content += line
	}
	data := []byte(content)
	err := os.WriteFile(FILENAME, data, 0644)
	check(err)
	fmt.Println("File has been created called", FILENAME)
}

//INFO: if error panic
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

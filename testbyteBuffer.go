package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	// Using string concatenation
	start := time.Now()
	str := ""
	for i := 0; i < 10000; i++ {
		str += "Hello"
	}
	fmt.Println("String concatenation took:", time.Since(start))

	// Using bytes.Buffer
	start = time.Now()
	var buf bytes.Buffer
	for i := 0; i < 10000; i++ {
		buf.WriteString("Hello")
	}
	fmt.Println("bytes.Buffer took:", time.Since(start))
}

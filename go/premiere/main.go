package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
    
    go spinner(1000)
    done := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
            if i % 10 == 0 {
                fmt.Println(i)
            }

            if(i == 50){
                done <- i
            }

            time.Sleep(100 * time.Millisecond)
		}
	}()

    v := <- done
    fmt.Println("Main goroutine")
    
    fmt.Println("What the fuck is",v)

    if v != 0 {
        os.Exit(v)
    }

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Println(fmt.Sprintf("\r%c", r))
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

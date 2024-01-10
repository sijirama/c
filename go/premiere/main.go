package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) append(data T) {
	stack.items = append(stack.items, data)
}

func (stack *Stack[T]) Pop() T {
	return stack.items[len(stack.items)-1]
}

func main() {

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

func worker(id int) {
	//mu.Lock()         // Get control of the conveyor belt
	//defer mu.Unlock() // Ensure the control is returned, even if something goes wrong

	fmt.Printf("Worker %d: Placing item on the conveyor belt.\n", id)
	counter++ // Modifying the shared resource (placing an item on the conveyor belt)
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Worker %d: Taking item from the conveyor belt.\n", id)
	counter-- // Modifying the shared resource (taking an item from the conveyor belt)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
    wg sync.WaitGroup
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

func (stack Stack[T]) Print() {
	for i := 0; i < len(stack.items); i++ {
		fmt.Println(stack.items[i])
	}
}

type Queue [T any] struct {
    items []T
    mux sync.RWMutex
}

func (q *Queue[T]) Enqueue (item T) {
    q.mux.Lock()
    defer q.mux.Unlock()
    q.items = append(q.items, item)
}

func (q *Queue[T]) Display (){
    for i := 0 ; i < len(q.items) ; i++ {
        fmt.Println(q.items[i])
    }
}

func inc[T any] (q *Queue[T] , i T){
    wg.Add(1)
    defer wg.Done()
    q.Enqueue(i)
}

func main() {
    q := Queue[int]{}
    go inc(&q,20)
    go inc(&q,40)
    go inc(&q,50)
    go inc(&q,60)
    go inc(&q,70)
    wg.Wait()
    q.Display()
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

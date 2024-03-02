package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mux   sync.Mutex
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.mux.Lock()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	wg.Done()
	c.mux.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()
	fmt.Printf("Result is %d", c.value)
}


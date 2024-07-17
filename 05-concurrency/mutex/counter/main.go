package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
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

	fmt.Printf("Result is %d\n", c.value)

}

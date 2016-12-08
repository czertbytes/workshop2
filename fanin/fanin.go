package main

import (
	"fmt"
	"sync"
)

// START OMIT
func fanIn(c1, c2 <-chan int) <-chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for next := range c1 {
			c <- next
		}
	}()
	go func() {
		defer wg.Done()
		for next := range c2 {
			c <- next
		}
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

// END OMIT

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	c := fanIn(c1, c2)

	go func() {
		for i := 0; i < 3; i++ {
			c1 <- i
			c2 <- i + 3
		}
		close(c1)
		close(c2)
	}()

	for next := range c {
		fmt.Printf("input is: %d\n", next)
	}
}

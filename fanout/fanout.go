package main

import (
	"fmt"
	"sync"
)

// START OMIT
func worker(id int, inCh <-chan int, wg *sync.WaitGroup) {
	for next := range inCh {
		fmt.Printf("%d: input is: %d\n", id, next)
	}
	wg.Done()
}

func main() {
	input := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, input, &wg)
	}
	for i := 0; i < 10; i++ {
		input <- i
	}
	close(input)
	wg.Wait()
}

// END OMIT

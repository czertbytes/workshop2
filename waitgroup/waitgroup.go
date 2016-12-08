package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}

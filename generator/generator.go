package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	for vi := range vowelsIndexGen("Hello Avocode") {
		fmt.Println(vi)
	}
}

func vowelsIndexGen(v string) <-chan int {
	ch := make(chan int)
	go func(v string) {
		for i, r := range v {
			switch strings.ToLower(string(r)) {
			case "a", "e", "i", "o", "u":
				ch <- i
			}
		}

		close(ch)
	}(v)
	return ch
}

// END OMIT

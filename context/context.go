package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func main() {
	defer func(start time.Time) {
		fmt.Printf("Took %s\n", time.Now().Sub(start))
	}(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if res, err := slowOperation(ctx); err == nil {
		fmt.Printf("Result is %d\n", res)
	}
}

func slowOperation(ctx context.Context) (int, error) {
	result := 0
	for i := 1; i < 6; i++ {
		time.Sleep(200 * time.Millisecond)
		result += i
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
		}
	}
	return result, nil
}

// END OMIT

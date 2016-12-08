package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

type article struct {
	URL, Title, Text string
	Words            []string
}

func main() {
	minScore := flag.Float64("min-score", 0.75, "The minimal score for similarity results.")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("\nUsage: articles [options]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	start := time.Now()

	var wg sync.WaitGroup
	resCh := make(chan result)
	quitCh := make(chan struct{})
	go func() {
		wg.Add(2)
		aArticles, bArticles := []article{}, []article{}
		go func() {
			defer wg.Done()
			for a := range idnesArticles() {
				aArticles = append(aArticles, a)
			}
		}()
		go func() {
			defer wg.Done()
			for a := range novinkyArticles() {
				bArticles = append(bArticles, a)
			}
		}()

		wg.Wait()

		wg.Add(len(aArticles) * len(bArticles))
		for _, a := range aArticles {
			for _, b := range bArticles {
				go func(a, b article) {
					defer wg.Done()
					score := cosineSimilarity(a.Words, b.Words)
					if score > *minScore {
						resCh <- result{
							A:     a,
							B:     b,
							Score: score,
						}
					}
				}(a, b)
			}
		}
		wg.Wait()
		close(quitCh)
	}()

	for {
		select {
		case r := <-resCh:
			fmt.Println(r)
		case <-quitCh:
			fmt.Printf("Took %s\n", time.Now().Sub(start))
			return
		}
	}
}

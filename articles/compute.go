package main

import (
	"fmt"
	"math"
)

type result struct {
	A, B  article
	Score float64
}

func (r result) String() string {
	return fmt.Sprintf("Score %f for articles %q (%s) and %q (%s)", r.Score, r.A.Title, r.A.URL, r.B.Title, r.B.URL)
}

func cosineSimilarity(txt1, txt2 []string) float64 {
	vect1 := make(map[string]int)
	for _, t := range txt1 {
		vect1[t]++
	}

	vect2 := make(map[string]int)
	for _, t := range txt2 {
		vect2[t]++
	}

	dotProduct := 0.0
	for k, v := range vect1 {
		dotProduct += float64(v) * float64(vect2[k])
	}

	sum1 := 0.0
	for _, v := range vect1 {
		sum1 += math.Pow(float64(v), 2)
	}

	sum2 := 0.0
	for _, v := range vect2 {
		sum2 += math.Pow(float64(v), 2)
	}

	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0.0
	}
	return float64(dotProduct) / float64(magnitude)
}

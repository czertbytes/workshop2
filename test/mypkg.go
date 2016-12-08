package test

func StringLen(s string) int {
	var counter int
	for range s {
		counter++
	}
	return counter
}

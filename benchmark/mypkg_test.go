package benchmark

// START OMIT

import (
	"strings"
	"testing"
)

func BenchmarkSHA512Hex(b *testing.B) {
	input := strings.Repeat("foo-bar-baz", 1000)
	for i := 0; i < b.N; i++ {
		SHA512Hex(input)
	}
}

// END OMIT

package benchmark

import (
	"crypto/sha512"
	"fmt"
	"io"
)

func SHA512Hex(s string) string {
	h := sha512.New()
	io.WriteString(h, s)
	return fmt.Sprintf("% x", h.Sum(nil))
}

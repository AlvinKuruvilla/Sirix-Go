package http

import (
	"bytes"
	"log"
)

//HeaderValue Represents an HTTP header field value
type HeaderValue struct {
	byteBuffer  bytes.Buffer
	isSensitive bool
}

//FromStatic Convert a static string to a `HeaderValue`
func FromStatic(s string) *HeaderValue {
	b := []byte(s)
	var buf bytes.Buffer
	for _, value := range b {
		if !isVisibleASCII(value) {
			log.Fatal("Invalid Header Value")
		}
		buf.Write(b)
	}
	return &HeaderValue{
		buf,
		false,
	}
}
func isVisibleASCII(b uint8) bool {
	// "|| b == b'\t'", this was after the 127, what does this mean?
	return (b >= 32 && b < 127)
}

// Exercise 4.7: in-place utf8 reverse
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello, world世!")
	reverse(s)
	fmt.Println(string(s))
}

func reverse(s []byte) {
	for i := 0; i < len(s); {
		// i is the start of the rest of the string to reverse.
		// Get the last rune in this string, and shift it to the start.
		r, size := utf8.DecodeLastRune(s[i:])
		copy(s[i+size:], s[i:len(s)-size])
		copy(s[i:i+size], []byte(string(r)))
		i += size
	}
}

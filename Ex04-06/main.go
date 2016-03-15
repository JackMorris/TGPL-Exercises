// Exercise 4.6: Convert runs of unicode spaces to a single ascii space
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello   world  !")
	s = spaceSquash(s)
	fmt.Println(string(s))
}

func spaceSquash(s []byte) []byte {
	bytesRemoved := 0
	for i := 0; i < len(s)-bytesRemoved; {
		r, size := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) || i+size >= len(s)-bytesRemoved {
			// Not a space, OR next character isn't valid.
			// Skip to the next character.
			i += size
			continue
		}

		// Get next character. If it's a space, squash down.
		// Otherwise, skip to the next character.
		r2, size2 := utf8.DecodeRune(s[i+size:])
		if unicode.IsSpace(r2) {
			copy(s[i+1:len(s)-(size+size2-1)], s[i+size+size2:])
			s[i] = ' '
			bytesRemoved += (size + size2 - 1)
			continue
		} else {
			i += size
		}
	}
	return s[:len(s)-bytesRemoved]
}

// Exercise 3.10: Non-recursive version of comma using bytes.Buffer
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	// Insert commas in `s` to separate collections of three characters.
	// e.g., "12345678" -> "12,345,678"
	// Returns the new string.
	if len(s) <= 3 {
		return s
	}

	var buffer bytes.Buffer
	var written = 0

	// Account for an 'odd' segment (eg. '12' in '12345678')
	var oddSegmentLength = len(s) % 3
	if oddSegmentLength > 0 {
		buffer.WriteString(s[:oddSegmentLength])
		buffer.WriteByte(',')
		written += oddSegmentLength
	}

	// Print remaining groups of 3.
	for ; written < len(s); written += 3 {
		buffer.WriteString(s[written : written+3])
		buffer.WriteByte(',')
	}

	// Ensure the final ',' is stripped off
	var ret = buffer.String()
	return ret[:len(ret)-1]
}

func main() {
	fmt.Println(comma("12"))       // "12"
	fmt.Println(comma("12345678")) // "12,345,678"
	fmt.Println(comma("123456"))   // "123,456"
}

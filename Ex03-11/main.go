// Exercise 3.11: Non-recursive version of comma using bytes.Buffer
// Also handles floating point values, and a sign.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	// Insert commas in `s` to separate collections of three characters.
	// e.g., "12345678" -> "12,345,678"
	// Returns the new string.

	var buffer bytes.Buffer

	// Identify any sign and floating point component.
	// This can be removed from the input string.
	// The sign can be added to the buffer now - the fpComponent
	// will be added later.
	sign := getSign(s)
	fpComponent := getFpComponent(s)
	buffer.WriteString(sign)
	s = s[len(sign) : len(s)-len(fpComponent)]

	var segmentLength = len(s) % 3 // Account for 'odd' segment (eg. "12" in "12345")
	for written := 0; written < len(s); {
		if segmentLength > 0 {
			buffer.WriteString(s[written : written+segmentLength])
			buffer.WriteByte(',')
		}
		written += segmentLength
		segmentLength = 3 // Standard segment length is 3.
	}

	// Ensure the final ',' is stripped off, and add back the fpComponent
	var commaString = buffer.String()
	return commaString[:len(commaString)-1] + fpComponent
}

func getSign(s string) string {
	// If s is a numerical string containing a sign character, it is returned.
	// E.g., "+123" -> "+", "456" -> ""
	if len(s) > 0 && (s[0] == '+' || s[1] == '-') {
		return string(s[0])
	}
	return ""
}

func getFpComponent(s string) string {
	// If s is a numerical string containing a floaring point component, it is returned
	// (along with the corresponding decimal point)
	// E.g., "123.456" -> ".456", "123" -> ""
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return s[dot:]
	}
	return ""
}

func main() {
	fmt.Println(comma("12"))            // "12"
	fmt.Println(comma("12345678"))      // "12,345,678"
	fmt.Println(comma("123456"))        // "123,456"
	fmt.Println(comma("12345.6754"))    // "12,345.6754"
	fmt.Println(comma("-45212.123456")) // "-45,212.123456
}

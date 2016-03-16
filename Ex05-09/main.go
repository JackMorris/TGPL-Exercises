// Exercise 5.9: functions as values
package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s := "Hello $Hello Test Test $Test"
	s = expand(s, reverse)
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		// If word start with a $ (eg. $foo), replace
		// with f(foo).
		if len(words[i]) > 0 && words[i][0] == '$' {
			words[i] = f(words[i][1:])
		}
	}
	return strings.Join(words, " ")
}

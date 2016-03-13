// Exercise 3.12: function to report whether two strings are anagrams of each other
package main

import (
	"fmt"
	"strings"
)

func isAnagram(s1, s2 string) bool {

	// Base case: two equal strings are obvious anagrams
	if s1 == s2 {
		return true
	}

	// Find if the first character of s1 is in s2.
	// If so, remove from both, and recurse
	// Otherwise, return false
	if len(s1) == len(s2) {
		if index := strings.Index(s2, string(s1[0])); index >= 0 {
			return isAnagram(s1[1:], s2[:index]+s2[index+1:])
		}
	}
	return false
}

func main() {
	fmt.Println(isAnagram("abc", "abc")) // True
	fmt.Println(isAnagram("abc", "cba")) // True
	fmt.Println(isAnagram("ab", "abc"))  // False
	fmt.Println(isAnagram("abc", "bcd")) // False
	fmt.Println(isAnagram("", ""))       // True
}

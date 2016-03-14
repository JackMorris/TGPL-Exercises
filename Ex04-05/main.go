// Exercise 4.5: in-place removal of adjacent duplicates
package main

import "fmt"

func main() {
	s := []string{"Hello", "World", "Hi", "Hi", "Hello"}
	s = removeAdjDuplicates(s)
	fmt.Println(s) // [Hello World Hi Hello]

	s = []string{"Hello", "World"}
	s = removeAdjDuplicates(s)
	fmt.Println(s) // [Hello World]

	s = []string{"Hi", "Hi", "Hi", "Hi"}
	s = removeAdjDuplicates(s)
	fmt.Println(s) // [Hi]
}

func removeAdjDuplicates(s []string) []string {
	duplicates := 0
	for i := 0; i < len(s)-1-duplicates; {
		if s[i] == s[i+1] {
			// Duplicate found - shift down.
			copy(s[i:len(s)-1], s[i+1:])
			duplicates++
		} else {
			i++
		}
	}
	return s[:len(s)-duplicates]
}

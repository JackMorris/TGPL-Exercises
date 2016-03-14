// Exercise 4.4: single pass rotate to the left
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s = rotate(s, 2)
	fmt.Println(s)
}

// rotate rotates s to the left by n
func rotate(s []int, n int) []int {
	len := len(s)
	ret := make([]int, len)
	for i := 0; i < len; i++ {
		ret[i] = s[(i+n)%len]
	}
	return ret
}

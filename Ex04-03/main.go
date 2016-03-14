// Exercise 4.3: function to reverse an array passed by address
package main

import "fmt"

const ARRAY_LEN = 10

func main() {
	a := [ARRAY_LEN]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&a)
	fmt.Println(a) // [10 9 8 7 6 5 4 3 2 1]
}

func reverse(a *[ARRAY_LEN]int) {
	for i, j := 0, ARRAY_LEN-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

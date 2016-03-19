// Exercise 5.15: Variadic min and max
package main

import "fmt"

func min(args ...int) (int, bool) {
	// Can't find minimum if no args are supplied.
	if len(args) == 0 {
		return 0, false
	}

	// Find minimum by looping over args.
	currentMin := args[0]
	for _, v := range args[1:] {
		if v < currentMin {
			currentMin = v
		}
	}
	return currentMin, true
}

func min2(v1 int, values ...int) int {
	currentMin := v1
	for _, v := range values {
		if v < currentMin {
			currentMin = v
		}
	}
	return currentMin
}

func max(args ...int) (int, bool) {
	// Can't find maximum if no args are supplied.
	if len(args) == 0 {
		return 0, false
	}

	// Find maximum by looping over args.
	currentMax := args[0]
	for _, v := range args[1:] {
		if v > currentMax {
			currentMax = v
		}
	}
	return currentMax, true
}

func max2(v1 int, values ...int) int {
	currentMax := v1
	for _, v := range values {
		if v > currentMax {
			currentMax = v
		}
	}
	return currentMax
}

func main() {
	vals := []int{1, 2, -5, 3, -9, 0}

	_, ok1 := min()
	fmt.Printf("min([]) OK? %t\n", ok1)

	minVal, ok2 := min(vals...)
	min2Val := min2(vals[0], vals[1:]...)
	fmt.Printf("min(%v) OK? %t\n", vals, ok2)
	fmt.Printf("min(%v) = %v\n", vals, minVal)
	fmt.Printf("min2(%v, %v) = %v\n", vals[0], vals[1:], min2Val)

	_, ok3 := max()
	fmt.Printf("max([]) OK? %t\n", ok3)

	maxVal, ok4 := max(vals...)
	max2Val := max2(vals[0], vals[1:]...)
	fmt.Printf("max(%v) OK? %t\n", vals, ok4)
	fmt.Printf("max(%v) = %v\n", vals, maxVal)
	fmt.Printf("max2(%v, %v), = %v\n", vals[0], vals[1:], max2Val)
}

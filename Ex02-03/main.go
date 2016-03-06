// Exercise 2.3: PopCount using loop.
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jackmorris/TGPL-Exercises/pkg/popcount"
)

func main() {
	rand.Seed(time.Now().Unix())
	val := uint64(rand.Int63())

	start := time.Now()
	popcount.PopCountLoop(val)
	dur := time.Since(start).Nanoseconds()
	fmt.Printf("PopCount Loop: %vns\n", dur)
}

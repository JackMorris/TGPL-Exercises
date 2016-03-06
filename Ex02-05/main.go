// Exercise 2.5: PopCount using quick check.
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
	popcount.PopCountQuick(val)
	dur := time.Since(start).Nanoseconds()
	fmt.Printf("PopCount Quick: %vns\n", dur)
}

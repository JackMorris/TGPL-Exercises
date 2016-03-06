// Exercise 2.4: PopCount using shift.
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
	popcount.PopCountShift(val)
	dur := time.Since(start).Nanoseconds()
	fmt.Printf("PopCount Shift: %vns\n", dur)
}

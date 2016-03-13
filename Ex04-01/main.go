// Exercise 4.1: count number of bits that are different in the SHA256 of the passed in messages
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: main.go msg1 msg2")
		return
	}

	msg1, msg2 := []byte(os.Args[1]), []byte(os.Args[2])
	hash1, hash2 := sha256.Sum256(msg1), sha256.Sum256(msg2)

	differences := 0
	for i := 0; i < 32; i++ {
		differences += byteDiffs(hash1[i], hash2[i])
	}
	fmt.Printf("Differences: %d\n", differences)
}

func byteDiffs(b1, b2 byte) int {
	// Return the number of bits that differ in b1 and b2

	// bitsDifferent contains a 1 bit if the corresponding bits of b1 and b2 are different
	bitsDifferent := b1 ^ b2

	// Count number of 1 bits in bitDiffs
	count := 0
	for ; bitsDifferent > 0; bitsDifferent &= bitsDifferent - 1 {
		count += 1
	}
	return count
}

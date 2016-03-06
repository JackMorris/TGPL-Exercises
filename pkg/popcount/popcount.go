// For Exercise 2.3, 2.4, 2.5.
// Package popcount counts the number of bits in a given uint64.
package popcount

// bc[i] = number of bits set in i, where i ranges 0->255.
var bc [256]byte

func init() {
	for i := range bc {
		// bc[i] = number of bits set for right-shifted i, plus 1 if i&1 != 0 (LSB is not 0).
		bc[i] = bc[i>>1] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		shiftFactor := uint(i * 8)
		count += int(bc[byte(x>>shiftFactor)])
	}
	return count
}

func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		shiftFactor := uint(i)
		count += int((x >> shiftFactor) & 1)
	}
	return count
}

func PopCountQuick(x uint64) int {
	count := 0
	for x&(x-1) != x {
		count++
		x &= (x - 1)
	}
	return count
}

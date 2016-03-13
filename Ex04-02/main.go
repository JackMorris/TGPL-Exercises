// Exercise 4.2: print the specified hash of the input
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var size = flag.Int("s", 256, "Variety of SHA to use (256, 384 or 512)")

func main() {
	flag.Parse()
	data, _ := ioutil.ReadAll(os.Stdin)

	switch *size {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(data))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(data))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(data))
	default:
		// Unsupported SHA varient
		flag.PrintDefaults()
	}
}

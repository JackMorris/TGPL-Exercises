// Exercises 1.2: Print index and value of each command line arg, on newlines
package main

import (
	"fmt"
	"os"
)

func main() {
	// Print program executable
	fmt.Println("P:" + os.Args[0])

	// Print each arg, with its index, starting at 1
	for index, arg := range os.Args[1:] {
		argNum := index + 1
		fmt.Print(argNum)
		fmt.Print(":" + arg + "\n")
	}
}

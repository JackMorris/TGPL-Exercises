// Exercise 1.4: from a list of files (names taken from command line args),
// determine the duplicate lines across these files.
// Prints such lines, along with the names of the files where the lines
// occur.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// occurances is a map from line -> (map of filenames -> count of that line in that file).
	occurances := make(map[string](map[string]int))

	// Build the occurances map.
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ex01-04: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			// Skip blank lines
			if line == "" {
				continue
			}

			// Make the innter map for that line if it doesn't already exist.
			if occurances[line] == nil {
				occurances[line] = make(map[string]int)
			}

			occurances[line][filename]++
		}
	}

	// For each line in the map, see if total count is > 1, and if so, print details
	// Also build a fileCountsDescription, used when printing the occurances per file, eg. "file1:2, file3:2".
	for line, countsPerFile := range occurances {
		totalCount := 0
		fileCountsDescription := ""
		for filename, count := range countsPerFile {
			totalCount += count
			fileCountsDescription += fmt.Sprintf("%s:%d, ", filename, count)
		}
		if totalCount > 1 {
			// Remeber to strip final two characters (', ') from fileCountsDescription.
			// This is fine, since totalCount > 1, so fileCountsDescription cannot be empty.
			fmt.Printf("%d\t%s (%s)\n", totalCount, line, fileCountsDescription[:len(fileCountsDescription)-2])
		}
	}
}

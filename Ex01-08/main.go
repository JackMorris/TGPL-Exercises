// Exercise 1.8: Print the content of each URL passed as an argument to stdout, adding http if not given.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		// Prepend protocal if not given.
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		// Fetch the contents of the URL into resp.
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Copy the body of resp into stdout.
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}

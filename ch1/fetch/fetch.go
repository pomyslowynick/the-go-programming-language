package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Exercises 1.7, 1.8, 1.9
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "Response status code: %d\n", resp.StatusCode)
		// b, err := ioutil.ReadAll(resp.Body)
		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}

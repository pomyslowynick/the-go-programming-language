// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Exercise 1.11
// Plenty of errors happening, I/O timeouts, TCP and TLS timeouts, handshake timeouts, request shutdowns.
// Madness.
// https://stackoverflow.com/questions/52017133/dial-tcp-i-o-timeout-on-simultaneous-requests#
func main() {
	f, err := os.Open("./top500Domains.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(bufio.NewReader(f))
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	ch := make(chan string, 3)
	file, err2 := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err2 != nil {
		log.Fatal(err)
	}

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		go fetch(line[1], ch) // start a goroutine
	}

	for i := 1; i < 500; i++ {
		if _, err := file.Write([]byte(<-ch)); err != nil {
			log.Fatal(err)
		}

	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get("http://" + url)
	if err != nil {
		ch <- fmt.Sprint(err, "\n") // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

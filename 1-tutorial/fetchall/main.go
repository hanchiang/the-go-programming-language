/*
* Fetchall fetches URLs in parallel and reports their times and sizes
* A goroutine is a concurrent function execution. A channel is a communication mechanism
* that allows one goroutine to pass values of a specified type to another goroutine. The function
* main runs in a goroutine and the go statement creates additional goroutines.
 */

// command: go run main.go gopl.io golang.org godoc.org
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // creates a channel of strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v", err) // send to channel ch
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d bytes %s", secs, nBytes, url)
}

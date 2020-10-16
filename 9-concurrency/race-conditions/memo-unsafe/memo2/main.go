package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hanchiang/the-go-programming-language/9-concurrency/race-conditions/httpGet"
	"github.com/hanchiang/the-go-programming-language/9-concurrency/race-conditions/memo-unsafe/memo"
)

// Make HTTP requests concurrently. Data race when fetching "http://www.google.com"
func main() {
	m := memo.New(httpGet.HttpGetBody)
	urls := []string{
		"http://www.google.com", "http://www.google.com", "http://www.facebook.com",
		"http://gopl.io"}
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()

}

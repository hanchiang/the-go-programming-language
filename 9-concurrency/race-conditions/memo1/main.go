package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hanchiang/the-go-programming-language/9-concurrency/race-conditions/httpGet"
	"github.com/hanchiang/the-go-programming-language/9-concurrency/race-conditions/memo"
)

func main() {
	m := memo.New(httpGet.HttpGetBody)
	urls := []string{
		"http://www.google.com", "http://www.google.com", "http://www.facebook.com",
		"http://gopl.io"}
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

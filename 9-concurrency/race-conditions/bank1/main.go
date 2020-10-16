package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		wg.Done()
	}()
	go func() {
		Deposit(100)
		fmt.Println("=", Balance())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("=", Balance())
}

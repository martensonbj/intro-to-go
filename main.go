package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked! But its fine. Message received: ", r)
	}
}

func printStuff() {
	// Decrement the wait group counter
	// Use defer so that if the function panics we aren't waiting forever
	// Also figure out what to do if the program panics
	defer wg.Done()
	defer handlePanic()

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// Increment the wait group counter
	wg.Add(1)
	// Launch a goroutine
	go printStuff()
	// Increment the wait group counter
	wg.Add(1)

	go printStuff()
	// Wait for the waitgroup counter to be 0 before continuing
	wg.Wait()
}

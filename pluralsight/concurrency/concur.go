package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// Allow processors to run two threads
	runtime.GOMAXPROCS(2)

	// Make main function wait for goroutines to finish before exiting
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		// Report when finished
		defer waitGroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		// Report when finished
		defer waitGroup.Done()

		fmt.Println("Pluralsight")
	}()

	// Wait until goroutines report they're finished
	waitGroup.Wait()
}

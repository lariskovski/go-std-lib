package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// init sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting. It is used
// to utilize all available CPU cores.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// foo prints numbers from 0 to 50 with a delay of 3 milliseconds
// between each print. It signals the WaitGroup when done.
func foo() {
	for i := 0; i <= 50; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

// bar prints characters from 'a' to 'z' with a delay of 20 milliseconds
// between each print. It signals the WaitGroup when done.
func bar() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("bar: %c\n", ch)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	// Add 2 to the WaitGroup counter to indicate that we are waiting
	// for two goroutines to finish.
	wg.Add(2)
	
	// Start foo and bar as goroutines to run them concurrently.
	go foo()
	go bar()
	
	// Wait for both goroutines to finish before exiting the main function.
	wg.Wait()
}
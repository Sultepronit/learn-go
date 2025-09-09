package main

// To wait for multiple goroutines to finish, we can use a wait group.
// SHOULD WORK WITH go1.25.0+!

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n")
}

func main() {
	var wg sync.WaitGroup
	fmt.Println(wg)

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Block until all the goroutines started by wg are done. A goroutine is done when the function it invokes returns.
	wg.Wait()
}
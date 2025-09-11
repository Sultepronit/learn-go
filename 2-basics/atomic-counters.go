package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// We’ll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64
	var simpleCounter int

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	// To atomically increment the counter we use Add.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
				simpleCounter++
			}
		})
	}

	// Wait until all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())
	fmt.Println("simple counter:", simpleCounter)
}

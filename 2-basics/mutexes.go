package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
type Container struct {
	mu             sync.Mutex
	counters       map[string]int
	simpleCounters map[string]int
}

// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// The zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters:       map[string]int{"a": 0, "b": 0},
		simpleCounters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	simpleIncrement := func(name string, n int) {
		for range n {
			c.simpleCounters[name]++
		}
	}

	// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	wg.Go(func() {
		doIncrement("a", 10000)
		simpleIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
		// simpleIncrement("a", 10000) // fatal error: concurrent map writes
	})

	wg.Go(func() {
		doIncrement("b", 10000)
		// simpleIncrement("b", 10000)
	})

	wg.Wait()
	fmt.Println(c.counters)       // map[a:20000 b:10000]
	fmt.Println(c.simpleCounters) // map[a:10000 b:0]
}

package main

// A goroutine is a lightweight thread of execution.

import (
	"fmt"
)

func f(from string) {
	// for i := range 3 {
	for i := range 100_000 {
		fmt.Println(from, ":", i)
		// time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	f("direct")

	go f("gorutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	f("direct 2")

	// without sleep, we'd have only direct output!
	// time.Sleep(time.Second)
	// but with i := range 100_000 it works perfectly fine!!!
}

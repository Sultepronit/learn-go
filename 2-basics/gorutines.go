package main

// A goroutine is a lightweight thread of execution.

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
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
	time.Sleep(time.Second)

}

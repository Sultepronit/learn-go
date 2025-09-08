package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

func worker(done chan bool) {
	fmt.Print("working... ")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

///////////////////// Channel Directions ///////////////////////
// When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg) // ping

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

	// Here we make a channel of strings buffering up to 2 values.
	msgs := make(chan string, 2)

	msgs <- "buffered"
	msgs <- "channel"

	fmt.Println(<-msgs) // buffered
	fmt.Println(<-msgs) // channel
	// fmt.Println(<-msgs) // fatal error: all goroutines are asleep - deadlock!

	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// If you removed the <- done line from this program, the program would exit before the worker even started.

	pings := make(chan string, 1)
	pongs := make(chan string, 2)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

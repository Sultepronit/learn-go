package main

import (
    "fmt"
    "time"
)

// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.

func main() {
    timer1 := time.NewTimer(1 * time.Second)

    <-timer1.C // blocks on the timer’s channel C until it sends a value indicating that the timer fired.
    fmt.Println("Timer 1 fired")

    // If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()

    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}
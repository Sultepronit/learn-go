package main

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request1", req, time.Now())
    }

    fmt.Println()

    // We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
    // burstyLimiter := make(chan time.Time, 3)
    burstyLimiter := make(chan any, 3)

    for range 3 {
        // burstyLimiter <- time.Now()
        burstyLimiter <- nil
    }

    go func() {
        // after the burstyLimiter buffer is filled up, the gorutine blocks
        // while the loop never breaks!
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t 
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request2", req, time.Now())
    }
}
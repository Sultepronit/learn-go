package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//  In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received.
			j, more := <-jobs
			if more {
				fmt.Println("received a job", j)
			} else {
				fmt.Println("all the jobs are done")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent a job")
	}
	close(jobs)
	fmt.Println("sent all the jobs")

	// We await the worker using the synchronization approach we saw earlier.
	<-done

	// Reading from a closed channel succeeds immediately, returning the zero value of the underlying type. The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero value generated because the channel is closed and empty.
	_, ok := <-jobs
	fmt.Println("there are more jobs:", ok)
}

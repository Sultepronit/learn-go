package main

import "fmt"

func main() {
	fmt.Println("hello there")

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	close(queue) // without: fatal error: all goroutines are asleep - deadlock!

	for elem := range queue {
		fmt.Println(elem)
	}
}
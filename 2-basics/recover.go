package main

import "fmt"

//  A recover can stop a panic from aborting the program and let it continue with execution instead.

func mayPanic() {
	panic("a problem!")
}

func main() {
	// recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:\n", r)
		}
	}()

	mayPanic()

	// This code will not run, because mayPanic panics. The execution of main stops at the point of the panic and resumes in the deferred closure.
	fmt.Println("expect nothing")
}

// recovered from:
//  a problem!

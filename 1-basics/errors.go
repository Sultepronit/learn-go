package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42!")
	}

	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boild water")

// We can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf.
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy some tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("It's dark!")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}

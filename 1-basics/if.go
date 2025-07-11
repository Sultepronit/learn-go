package main

import (
	"fmt"
	"time"
)

func main() {
	if 2 < 25 {
		fmt.Println("small")
	} else {
		fmt.Println("big")
	}

	if false || 5 > 1 {
		fmt.Println("true!")
	}

	if num := 9; num < 0 {
		fmt.Println("negative!")
	} else if num > 1000 {
		fmt.Println("Big!")
	} else {
		fmt.Println(num)
	}

	// fmt.Println(num) // doesn't exists here!

	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println(time.Monday)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // "a type switch"; i.(type) have no sense outside?..
		case bool:
			fmt.Println(t, "- bool!")
		case int:
			fmt.Println(t, " - int!")
		default:
			fmt.Println(t)
		}
	}

	whatAmI(true)
	whatAmI(14)
	whatAmI("hello!")
}

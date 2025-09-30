package main

import (
	"fmt"
	"os"
)

// Go offers excellent support for string formatting in the printf tradition.

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}
	// Go offers several printing “verbs” designed to format general Go values.
	pf("%v\n", p)  // {1 2}
	pf("%+v\n", p) // {x:1 y:2}
	pf("%#v\n", p) // main.point{x:1, y:2}

	pf("type: %T\n", p)     // type: main.point
	pf("type: %T\n", 589)   // type: int
	pf("type: %T\n", 589.0) // type: float64

	s := "\"string\"\tA"
	pf("string: %s\n", s)      // string: "string"        A
	pf("string: %q\n", s)      // string: string: "\"string\"\tA"
	pf("bool: %t\n", true)     // bool: true
	pf("pointer: %p\n", &p)    // pointer: 0xc000088020
	pf("int (deс): %d\n", 598) // int (deс): 598
	pf("bin: %b\n", 15)        // bin: 1111
	pf("hex: %x\n", 100)       // hex: 64
	pf("float: %f\n", 3.14)    // float: 3.140000
	pf("sci: %e\n", 3.14)      // sci: 3.140000e+00
	pf("sci: %E\n", 3.14)      // sci: 3.140000E+00

	// %x renders the string in base-16, with two output characters per byte of input.
	pf("string->hex: %x\n", s) // 22737472696e67220941
	pf("char: %c\n", 33)       // char: !

	pf("|%6d|%6d|\n", 11, 256)            // |    11|   256|
	pf("|%6.2f|%6.2f|\n", 10.0, 13.256)   // | 10.00| 13.26|
	pf("|%-6d|%-6.2f|\n", 7, 3.145)       // |7     |3.15  |
	pf("|%6s|%6s|\n", "hey world", "app") // |hey world|   app|
	pf("|%-6s|%-6s|\n", "apple", "pie")   // |apple |pie   |

	str := fmt.Sprintf("%9s", "go!")
	fmt.Println(str) //       go!

	fmt.Fprintf(os.Stderr, "Error: %s\n", "Something went wrong!") // Error: Something went wrong!
}

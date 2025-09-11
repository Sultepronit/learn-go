package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.

var pl = fmt.Println

func createFile(path string) *os.File {
	pl("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File, data string) {
	pl("writing")
	fmt.Fprintln(file, data)
}

func main() {

}

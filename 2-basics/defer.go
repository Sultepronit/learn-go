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

func closeFile(file *os.File) {
	pl("closing")
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Immediately after getting a file object with createFile, we defer the closing of that file with closeFile. This will be executed at the end of the enclosing function (main), after writeFile has finished.
	f := createFile("/tmp/file.txt")
	defer closeFile(f)
	writeFile(f, "Hello there!")
}

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pl(match) // true

	// Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")
	pl(r.MatchString("peach"))                          // true
	pl(r.Match([]byte("peach")))                        // true
	pl(r.FindString("peach punch"))                     // peach
	pl(r.FindStringIndex("peach punch"))                // [0 5]
	pl(r.FindStringSubmatch("peach punch"))             // [peach ea] results for p([a-z]+)ch & ([a-z]+)
	pl(r.FindStringSubmatchIndex("peach punch"))        // [0 5 1 3] results for p([a-z]+)ch & ([a-z]+)
	pl(r.FindAllString("peach punch pinch", -1))        // [peach punch pinch]
	pl(r.FindAllString("peach punch pinch", 2))         // [peach punch]
	pl(r.FindAllStringSubmatchIndex("peach punch", -1)) // [[0 5 1 3] [6 11 7 9]]
	pl(r.ReplaceAllString("a peach", "fruit"))          // a fruit

	re := r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)
	pl(string(re)) // a PEACH

	// When creating global variables with regular expressions you can use the MustCompile variation of Compile. MustCompile panics instead of returning an error, which makes it safer to use for global variables.
	r = regexp.MustCompile("p([a-z]+)ch")
	pl(r) // p([a-z]+)ch
}

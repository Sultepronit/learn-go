package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains:", s.Contains("test", "es"))         // contains: true
	p("count:", s.Count("test", "t"))                // count: 2
	p("hasPrefix:", s.HasPrefix("test", "te"))       // hasPrefix: true
	p("hasSuffix:", s.HasSuffix("test", "st"))       // hasSuffix: true
	p("index:", s.Index("test", "e"))                // 1
	p("join:", s.Join([]string{"a", "b"}, "."))      // join: a.b
	p("repeat:", s.Repeat("F", 5))                   // repeat: FFFFF
	p("replace:", s.Replace("banana", "a", "A", 2))  // replace: bAnAna
	p("replace:", s.Replace("banana", "a", "A", -1)) // replace: bAnAnA
	p("replace:", s.Replace("abc", "", "-", -1))     // replace: -a-b-c-
	p("stplit:", s.Split("a-b-c-d-e", "-"))          // stplit: [a b c d e]
	p("lower:", s.ToLower("TEST"))                   // lower: test
	p("upper:", s.ToUpper("test"))                   // upper: TEST
}

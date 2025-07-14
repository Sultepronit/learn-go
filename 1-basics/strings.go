package main

import (
	"fmt"
	"unicode/utf8"
)

// https://go.dev/blog/strings
// A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8.

func main() {
	s := "abcあいう"

	fmt.Println(len(s)) // 12
	// fmt.Println(cap(s)) // ok, it doesn't work

	fmt.Println(utf8.RuneCountInString(s)) // 6

	for _, c := range s {
		// fmt.Println(i, c)
		fmt.Printf("%d -> %x -> %#U\n", c, c, c)
	}
	// 97 -> 61 -> U+0061 'a'
	// ...
	// 12358 -> 3046 -> U+3046 'う'

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
		// fmt.Printf("%x\n", s[i])
	}

	kana := "あいう"

	rune, size := utf8.DecodeRuneInString(kana)
	fmt.Println(rune, size) // 12354 3

	for i := 0; i < len(kana); i++ {
		rune, size := utf8.DecodeRuneInString(kana[i:])
		fmt.Println(rune, size)
		fmt.Printf("%#U\n", rune)
		if rune == 'い' { // 'い' - rune literal
			fmt.Println("い found!")
			break
		}
	}
}

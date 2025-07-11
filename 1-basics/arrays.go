package main

import (
	"fmt"
	"slices"
)

// https://go.dev/blog/slices-intro
func main() {
	// arrays -------------------------------------------------
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a))

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// b = [...]int{7, 8, 9}
	// cannot use [...]int{…} (value of type [3]int) as [5]int value in assignment

	b = [...]int{7, 8, 9, 10, 11}
	fmt.Println(b)

	// b = [...]int{100, 2: 400, 500}
	// cannot use [...]int{…} (value of type [4]int) as [5]int value in assignment

	b = [...]int{100, 3: 400, 500}
	fmt.Println(b) // [100 0 0 400 500]

	c := [...]int{100, 9: 200, 2}
	fmt.Println(c) // [100 0 0 0 0 0 0 0 0 200 2]

	var d [2][3]int
	for i := range 2 {
		for j := range 3 {
			d[i][j] = i + j
		}
	}
	fmt.Println(d) // [[0 1 2] [1 2 3]]

	// e := [2][3]int{
	e := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(e)

	// slices -------------------------------------------------
	var s []string
	fmt.Println(s, s == nil, len(s)) // [] true 0

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s)) // [  ] 3 3

	s[0] = "a"
	s[1] = "b"
	s = append(s, "d")
	fmt.Println(s, len(s), cap(s)) // [a b  d] 4 6

	s2 := append(s, "e", "f")
	fmt.Println(s2, len(s2), cap(s2)) // [a b  d e f] 6 6

	n := make([]int, 5)
	n[3] = 77
	n2 := append(n, 1, 100)
	fmt.Println(n2, len(n2), cap(n2)) // [0 0 0 77 0 1 100] 7 10
	// fmt.Println(n2[9]) // panic: runtime error: index out of range [9] with length 7

	cs := make([]string, len(s))
	copy(cs, s)
	fmt.Println(cs) // [a b  d]

	fmt.Println(s[1:])  // [b  d] "slice" operator: slices from 1, including
	fmt.Println(s[:3])  // [a b ] slices up to 3, excluding
	fmt.Println(s[1:3]) // [b ] slices from 1 to 2
	fmt.Println(s[1:6]) // [b  d e f] it... uses s2!
	// fmt.Println(s[1:7]) // slice bounds out of range [:7] with capacity 6 - yeah s2 cap = 6

	se := []int{1, 2, 3}
	se2 := append(se, 4)        // new underlying array
	se3 := append(se2[1:2], 99) // same underlying array

	fmt.Println("se:", se)   // se: [1 2 3]
	fmt.Println("se2:", se2) // se2: [1 2 99 4]
	fmt.Println("se3:", se3) // se3: [2 99]

	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}
	// fmt.Println(t == t2) // invalid operation: t == t2 (slice can only be compared to nil)
	fmt.Println(slices.Equal(t, t2)) // true

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	tds := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		tds[i] = make([]int, innerLen)
		for j := range innerLen {
			tds[i][j] = i + j
		}
	}
	fmt.Println(tds) // [[0] [1 2] [2 3 4]]
}

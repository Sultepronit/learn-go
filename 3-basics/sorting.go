package main

import (
	"cmp"
	"fmt"
	"slices"
)

var p = fmt.Println

func simpleSort() {
	strs := []string{"z", "b", "a", "f"}
	slices.Sort(strs)
	p(strs) // [a b f z]

	ints := []int{5, 8, 1}
	slices.Sort(ints)
	p(ints) // [1 5 8]

	p("is sorted:", slices.IsSorted(ints)) // is sorted: true
}

func customSort() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenComp)
	p(fruits) // [kiwi peach banana]

	type person struct {
		name string
		age  uint8
	}

	people := []person{
		{"Stepan", 32},
		{"Anna", 17},
		{"Ichiro", 25},
	}

	slices.SortFunc(people,
		func(a, b person) int {
			return cmp.Compare(a.age, b.age)
		})
	p(people) // [{Anna 17} {Ichiro 25} {Stepan 32}]
}

func main() {
	simpleSort()
	customSort()
}

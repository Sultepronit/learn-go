package main

import "fmt"

// https://gobyexample.com/generics

// As an example of a generic function, SlicesIndex takes a slice of any comparable type and an element of that type and returns the index of the first occurrence of v in s, or -1 if not present. The comparable constraint means that we can compare values of this type with the == and != operators.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// As an example of a generic type, List is a singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"a", "b", "C"}
	fmt.Println(SlicesIndex(s, "b")) // 1
	// we can specyfy types:
	fmt.Println(SlicesIndex[[]string, string](s, "C")) // 2

	lst := List[int]{}
	lst.Push(3)
	lst.Push(2)
	lst.Push(1)
	fmt.Println(lst.AllElements()) // [3 2 1]
}

package main

import "fmt"

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

type person struct {
	name string
	age  int
}

// It’s idiomatic to encapsulate new struct creation in constructor functions
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// methods
type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type test struct {
	a, b int
}

func (t *test) setA(a int) {
	t.a = a
}

// does nothing - works on copy!!!
func (t test) setB(b int) {
	t.b = b
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	sean := person{"Sean", 50}
	fmt.Println(sean.name)

	pointer := &sean
	fmt.Println(pointer.age) // 50

	sean.age = 45
	fmt.Println(pointer.age) //

	// an anonymous struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) // {Rex true}

	r := rect{10, 5}
	fmt.Println("area:", r.area())   // area: 50
	fmt.Println("perim:", r.perim()) // perim: 30

	t := test{}
	fmt.Println(t) // {0 0}
	t.setA(11)
	t.setB(22)
	fmt.Println(t) // {11 0}
}

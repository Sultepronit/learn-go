package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {
	c1 := container{}
	fmt.Println(c1) // {{0} }

	c2 := container{
		base: base{
			num: 2,
		},
		str: "some text",
	}
	fmt.Println(c2) // {{2} some text}

	// we can spell out the full path using the embedded type name,
	// or use it's fields directly
	fmt.Println(c2.base.num)   // 2
	fmt.Println(c2.num)        // 2
	fmt.Println(c2.describe()) // base with num = 2

	// yeah, container has describe() method (it embeds base),
	// so it implements the describer interface!!!
	type describer interface {
		describe() string
	}

	var d describer = c2
	fmt.Println(d.describe()) // base with num = 2
}

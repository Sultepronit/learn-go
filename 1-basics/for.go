package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 10; i < 13; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop!")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	fmt.Println("ranges:")
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	map1 := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go!" {
		fmt.Println(i, c)
	}
	// 0 103
	// 1 111
	// 2 33
}

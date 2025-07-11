package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	fmt.Println(m) // map[]

	m["a1"] = 12
	m["b7"] = 55
	fmt.Println(m) // map[a1:12 b7:55]

	fmt.Println(m["a1"]) // 12
	fmt.Println(m["a7"]) // 0 - If the key doesn’t exist, the zero value of the value type is returned.
	v, prs := m["a7"]
	fmt.Println(v, prs) // 0 false - The optional second return value when getting a value from a map indicates if the key was present in the map.

	fmt.Println(len(m)) // 2

	delete(m, "a1")
	fmt.Println(m) // map[b7:55]

	clear(m)
	fmt.Println(m) // map[]

	n := map[int16]string{1: "one", 2: "two"}
	n2 := map[int16]string{1: "one", 2: "two"}
	fmt.Println(maps.Equal(n, n2)) // true

	n3 := map[int]string{1: "one", 2: "two"}
	// fmt.Println(maps.Equal(n, n3)) // in call to maps.Equal, M2 (type map[int]string) does not satisfy ~map[K]V
}

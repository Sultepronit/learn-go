package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func ps(data []byte) {
	fmt.Println(string(data))
}

func main() {
	j1, _ := json.Marshal(true)
	ps(j1) // true

	j2, _ := json.Marshal([]string{"apple", "banana", "pear"})
	ps(j2) // ["apple","banana","pear"]

	j3, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	ps(j3) // {"apple":5,"lettuce":7}

	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "banana", "pear"},
	}
	j4, _ := json.Marshal(res1)
	ps(j4) // {"Page":1,"Fruits":["apple","banana","pear"]}

	res2 := &response2{2, []string{"apple", "banana", "pear"}}
	j5, _ := json.Marshal(res2)
	ps(j5) // {"page":2,"fruits":["apple","banana","pear"]}
}

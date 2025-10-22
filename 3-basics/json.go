package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var pl = fmt.Println

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

	jb := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(jb, &dat); err != nil {
		panic(err)
	}
	pl(dat) // map[num:6.13 strs:[a b]]

	// In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in num to the expected float64 type.
	num := dat["num"].(float64)
	pl(num) // 6.13

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	pl(str1) // a

	// We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	pl(res) // {1 [apple peach]}
	pl(res.Fruits[0]) // apple

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}

	dec := json.NewDecoder(strings.NewReader(str))
	res21 := response2{}
	dec.Decode(&res21)
	pl(res21) // {1 [apple peach]}
}

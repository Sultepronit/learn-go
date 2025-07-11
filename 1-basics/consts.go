package main

import (
	"fmt"
	"math"
)

const s string = "constant!"

func main() {
	fmt.Println(s)

	const n = 5000000
	const d = 3e20 / n
	fmt.Println(n, d)

	fmt.Println(int64(d))
	fmt.Println(float64(n))

	fmt.Println(math.Cos(n))

	fmt.Println(int16(math.Cos(n) * 50))
}

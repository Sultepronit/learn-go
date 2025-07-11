package main

import (
	"fmt"
	"log"
)

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multipleReturn() (int, string) {
	return 0, "we did it"
}

func variadicFunction(nums ...int) {
	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println("frst:", nums[0])

	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)
}

// anonymus fns & closures
// "func() int" is the return type!
func intSeq() func() int {
	i := 0
	// The returned function closes over the variable i to form a closure.
	return func() int {
		i++
		return i
	}
}

func main() {
	log.Println(plusPlus(10, 200, 1000))

	n, s := multipleReturn() // multiple assignment
	fmt.Println(n, s)

	variadicFunction(1, 20, 300, 4000)
	nums := []int{7, 6, 5, 4}
	variadicFunction(nums...)

	counter1 := intSeq()
	counter2 := intSeq()

	fmt.Println(counter1()) // 1
	fmt.Println(counter2()) // 1
	fmt.Println(counter1()) // 2

}

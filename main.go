package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/solutions"
)

func main() {
	//input := []int{0, 1, 0, 0, 2, 0, 0, 1}
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(input)
	fmt.Println(solutions.Trap(input))
}

package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/maxPointsOnALine"
)

func main() {
	inputs := [][][]int{
		{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
		{{1, 1}, {1, 1}, {2, 2}, {2, 2}},
	}

	for _, input := range inputs {
		fmt.Println(maxPointsOnALine.MaxPointsOnALine(input))
	}

}

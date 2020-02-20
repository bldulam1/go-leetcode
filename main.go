package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/maxPointsOnALine"
)

func main() {
	inputs := [][][]int{
		//{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
		//{{1, 1}, {1, 1}, {2, 2}, {2, 2}, {3, 3}},
		{{1, 1}, {1, 1}, {1, 1}},
		//{{0, 0}, {1, 1}, {1, -1}},
		//{{-4,-4},{-8,-582},{-3,3},{-9,-651},{9,591}},
		{{84, 250}, {0, 0}, {1, 0}, {0, -70}, {0, -70}, {1, -1}, {21, 10}, {42, 90}, {-42, -230}},
	}

	for _, input := range inputs {
		fmt.Println(maxPointsOnALine.MaxPoints(input))
	}

}

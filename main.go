package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/solutions"
)

func main() {
	ll := solutions.LadderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(ll)
}

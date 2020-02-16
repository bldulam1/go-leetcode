package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/validNumber"
)

func main() {
	testStrings := []string{
		"0123456789 e - +",
		"0.2",
		"-0.2",
		"1e4",
		"1e4e",
		"53.5e93",
		"-+3",
		"  .1",
		"1e",
		".e1",
		"te1",
		"6+1",
		" -.",
		"+",
	}

	for _, testString := range testStrings {
		fmt.Println(testString, validNumber.IsNumber(testString))
	}
}

package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/longestPalindrome"
)

func main() {
	//fmt.Println(medianOfTwoSortedArrays.MedianOfTwoSortedArrays([]int{2}, []int{1, 3}))
	fmt.Println(longestPalindrome.LongestPalindrome("babad"))
	fmt.Println(longestPalindrome.LongestPalindrome("a"))
	fmt.Println(longestPalindrome.LongestPalindrome("cbbd"))
}

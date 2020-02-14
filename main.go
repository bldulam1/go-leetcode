package main

import (
	"fmt"
	"github.com/bldulam1/go-leetcode/add_two_numbers"
)

func main() {
	l1 := &add_two_numbers.ListNode{
		Val: 5,
	}

	l2 := &add_two_numbers.ListNode{
		Val: 5,
	}

	sumList := add_two_numbers.AddTwoNumbers(l1, l2)
	fmt.Println(sumList)
}

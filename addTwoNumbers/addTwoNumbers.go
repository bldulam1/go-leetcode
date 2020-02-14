package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, current *ListNode
	n1, n2 := l1, l2

	carryOver := 0

	for n1 != nil || n2 != nil {
		var n1Val, n2Val int
		if n1 != nil {
			n1Val = n1.Val
		}
		if n2 != nil {
			n2Val = n2.Val
		}

		tempSum := n1Val + n2Val + carryOver
		carryOver = tempSum / 10

		if current != nil {
			current.Next = &ListNode{Val: tempSum % 10}
			current = current.Next
		} else {
			head = &ListNode{Val: tempSum % 10}
			current = head
		}

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	if carryOver > 0 && current != nil {
		current.Next = &ListNode{Val: carryOver}
	}

	return head
}

package add_two_numbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(node *ListNode) {
	cur := node
	for cur != nil {
		fmt.Print(cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentL1 := l1
	currentL2 := l2

	result := &ListNode{Val: 0, Next: nil}
	current := result
	sum := 0
	carry := 0

	for {

		if currentL1 == nil || currentL2 == nil {
			break
		}
		// both has digit
		sum = currentL1.Val + currentL2.Val + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		current.Val = sum
		current.Next = &ListNode{Val: 0, Next: nil}
		current = current.Next
		currentL1 = currentL1.Next
		currentL2 = currentL2.Next
	}

	for currentL1 != nil {
		sum = currentL1.Val + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		current.Val = sum
		current.Next = &ListNode{Val: 0, Next: nil}
		current = current.Next
		currentL1 = currentL1.Next
	}

	for currentL2 != nil {
		sum = currentL2.Val + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		current.Val = sum
		current.Next = &ListNode{Val: 0, Next: nil}
		current = current.Next
		currentL2 = currentL2.Next
	}

	if carry != 0 {
		current.Val = carry
	}

	current = result
	for current != nil && current.Next != nil {
		if current.Next.Next == nil && current.Next.Val == 0 {
			current.Next = nil
			break
		}
		current = current.Next
	}
	return result
}

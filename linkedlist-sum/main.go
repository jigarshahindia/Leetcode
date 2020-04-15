package main

import (
	"fmt"
)

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(number1 *ListNode, number2 *ListNode) *ListNode {
	add := number1.Val + number2.Val
	quotient := add / 10
	sum := add % 10
	number := ListNode{
		sum,
		nil,
	}
	previous := &number
	number1 = number1.Next
	number2 = number2.Next
	for number1 != nil && number2 != nil {
		add = quotient + number1.Val + number2.Val
		quotient = add / 10
		sum = add % 10
		number1 = number1.Next
		number2 = number2.Next
		previous.Next = &ListNode{
			sum,
			nil,
		}
		previous = previous.Next
	}
	for number1 != nil {
		add = quotient + number1.Val
		quotient = add / 10
		sum = add % 10
		number1 = number1.Next
		previous.Next = &ListNode{
			sum,
			nil,
		}
		previous = previous.Next
	}
	for number2 != nil {
		add = quotient + number2.Val
		quotient = add / 10
		sum = add % 10
		number2 = number2.Next
		previous.Next = &ListNode{
			sum,
			nil,
		}
		previous = previous.Next
	}
	if quotient != 0 {
		previous.Next = &ListNode{
			quotient,
			nil,
		}
	}
	return &number
}

func main() {
	number1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	number2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(&number1, &number2))
}

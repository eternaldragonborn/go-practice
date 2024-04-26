package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(data []int) *ListNode {
	var head, p *ListNode

	for _, x := range data {
		if head == nil {
			head = &ListNode{x, nil}
			p = head
		} else {
			p.Next = &ListNode{x, nil}
			p = p.Next
		}
	}

	return head
}

func (head *ListNode) ToSlice() []int {
	result := []int{}

	for p := head; p != nil; p = p.Next {
		result = append(result, p.Val)
	}

	return result
}

func (head *ListNode) Print() string {
	result := ""

	for p := head; p != nil; p = p.Next {
		result += fmt.Sprintf("%d", p.Val)

		if p.Next != nil {
			result += " "
		}
	}

	return result
}

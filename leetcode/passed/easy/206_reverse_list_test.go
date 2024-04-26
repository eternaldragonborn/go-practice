package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/list"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// nodes := []*ListNode{}
	var prev, p *ListNode = nil, head

	for p != nil {
		// nodes = append(nodes, p)
		next := p.Next
		p.Next = prev
		prev, p = p, next
		// p = next
	}

	// head = nodes[len(nodes)-1]
	// for i := len(nodes) - 1; i > 0; i-- {
	// 	nodes[i].Next = nodes[i-1]
	// }
	// nodes[0].Next = nil

	// return prev
	head = prev
	return head
}

func TestListReverse(t *testing.T) {
	cases := []struct {
		name string
		nums []int
	}{
		{"long", []int{1, 2, 3, 4, 5}},
		{"short", []int{1, 2}},
		{"one element", []int{1}},
		{"empty", []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			head := NewList(c.nums)
			ans := reverseList(head).ToSlice()
			slices.Reverse(c.nums)

			assert.Equal(t, c.nums, ans)
		})
	}
}

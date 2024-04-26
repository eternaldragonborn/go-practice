package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/list"
)

func hasCycle(head *ListNode) bool {
	// no any node or only one node
	if head == nil || head.Next == nil {
		return false
	}

	p1, p2 := head, head

	for p1 != nil && p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}

type LinkedListCycleCase struct {
	List []int
	Pos  int
	// Excepted bool
}

func createLinkedListCycleCase(t *testing.T, nums []int, pos int) *ListNode {
	t.Helper()
	var head, p, start *ListNode

	for index, n := range nums {
		node := new(ListNode)
		node.Val = n

		if head == nil {
			head = node
			p = head
		} else {
			p.Next = node
			p = p.Next
		}

		if index == pos {
			start = node
		}
	}

	if pos != -1 {
		p.Next = start
	}

	return head
}

func TestHasCycle(t *testing.T) {
	cases := []LinkedListCycleCase{
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 2},
		{[]int{3, 2, 0, -4}, 1},
		{[]int{1, 2, 3, 4, 2, 3, 2}, 1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{1, 2, 3}, -1},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, -1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			excepted := c.Pos != -1
			head := createLinkedListCycleCase(t, c.List, c.Pos)
			assert.Equal(t, excepted, hasCycle(head))
		})
	}
}

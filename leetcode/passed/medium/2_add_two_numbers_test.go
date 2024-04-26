package leetcode

import (
	"fmt"
	"testing"

	. "leetcode.com/fafnir/structure/list"
)

func isOverflow(val int, overflow bool) (int, bool) {
	if overflow {
		val++
	}

	if val >= 10 {
		val %= 10
		overflow = true
	} else {
		overflow = false
	}

	return val, overflow
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, p *ListNode
	p1, p2 := l1, l2
	overflow := false

	// both list pointer are not nil
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		newNode := new(ListNode)
		newNode.Val, overflow = isOverflow(p1.Val+p2.Val, overflow)

		if result == nil {
			result, p = newNode, newNode
		} else {
			p.Next = newNode
			p = p.Next
		}
	}

	// p1 is not nil, but p2 is
	for ; p1 != nil; p1 = p1.Next {
		newNode := new(ListNode)
		newNode.Val, overflow = isOverflow(p1.Val, overflow)

		p.Next = newNode
		p = p.Next
	}
	// p2 is not nil
	for ; p2 != nil; p2 = p2.Next {
		newNode := new(ListNode)
		newNode.Val, overflow = isOverflow(p2.Val, overflow)

		p.Next = newNode
		p = p.Next
	}

	if overflow {
		p.Next = new(ListNode)
		p.Next.Val = 1
	}

	return result
}

func TestAddNumbers(t *testing.T) {
	cases := []struct {
		L1, L2   []int
		Excepted string
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, "7 0 8"},
		{[]int{0}, []int{0}, "0"},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, "8 9 9 9 0 0 0 1"},
		{[]int{2, 4, 9}, []int{5, 6, 4, 9}, "7 0 4 0 1"},
	}

	for i, c := range cases {
		l1, l2 := NewList(c.L1), NewList(c.L2)

		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			if ans := addTwoNumbers(l1, l2).Print(); ans != c.Excepted {
				t.Logf("l1: %v, l2: %v", c.L1, c.L2)
				t.Errorf("excepted %s, got %s", c.Excepted, ans)
			}
		})
	}
}

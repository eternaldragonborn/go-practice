package leetcode

import (
	"slices"
	"sort"
	"testing"

	. "leetcode.com/fafnir/structure/list"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	var result, p *ListNode

	// first node
	if p1 == nil && p2 == nil {
		return nil
	}
	if p2 == nil || p1 != nil && p1.Val < p2.Val {
		result = p1
		p1 = p1.Next
	} else {
		result = p2
		p2 = p2.Next
	}
	p = result

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
	}

	return result
}

func TestNewList(t *testing.T) {
	cases := []struct {
		Data []int
	}{
		{[]int{1, 3, 2}},
		{[]int{1}},
	}

	t.Run("new list node", func(t *testing.T) {
		for _, c := range cases {
			s := NewList(c.Data).ToSlice()
			if !slices.Equal(s, c.Data) {
				t.Fatalf("excepted %v, got %v", s, c.Data)
			}
		}
	})
}

func TestMergeList(t *testing.T) {
	cases := []struct {
		List1 []int
		List2 []int
		// Excepted string
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{}, []int{}},
		{[]int{}, []int{0}},
		{[]int{0}, []int{}},
	}

	t.Run("test sort and merge list", func(t *testing.T) {
		for _, c := range cases {
			excepted := append(c.List1, c.List2...)
			sort.Slice(excepted, func(i, j int) bool { return excepted[i] < excepted[j] })

			newList := mergeTwoLists(NewList(c.List1), NewList(c.List2)).ToSlice()

			if !slices.Equal(newList, excepted) {
				t.Fatalf("excepted %v, got %v", excepted, newList)
			} else {
				t.Logf("result: %v", newList)
			}
		}
	})
}

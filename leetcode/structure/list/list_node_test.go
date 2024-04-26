package listnode

import (
	"slices"
	"testing"
)

func TestNewList(t *testing.T) {
	cases := []struct {
		Data     []int
		Excepted string
	}{
		{[]int{2, 4, 3}, "2 4 3"},
		{[]int{}, ""},
		{[]int{2}, "2"},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			list := NewList(c.Data)
			if ans := list.Print(); ans != c.Excepted {
				t.Errorf("excepted %s, got %s", c.Excepted, ans)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	cases := [][]int{
		{2, 4, 3},
		{},
		{2},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			list := NewList(c)
			if ans := list.ToSlice(); !slices.Equal(ans, c) {
				t.Errorf("excepted %v, got %v", c, ans)
			}
		})
	}
}

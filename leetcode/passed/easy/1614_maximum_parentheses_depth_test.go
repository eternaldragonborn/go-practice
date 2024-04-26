package leetcode

import (
	"testing"
)

func maxDepth(s string) int {
	depth, maxDepth := 0, 0

	for _, c := range s {
		if c == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if c == ')' {
			depth--
		}
	}

	return maxDepth
}

func TestMaxDepth(t *testing.T) {
	cases := []struct {
		S        string
		Excepted int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"()", 1},
		{"(()())", 2},
	}

	for _, c := range cases {
		t.Run(c.S, func(t *testing.T) {
			if ans := maxDepth(c.S); ans != c.Excepted {
				t.Errorf("excepted %d, got %d", c.Excepted, ans)
			}
		})
	}
}

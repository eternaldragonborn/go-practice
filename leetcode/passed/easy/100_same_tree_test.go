package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) ||
		(q == nil && p != nil) {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestIsSameTree(t *testing.T) {
	cases := []struct {
		P, Q []any
	}{
		{[]any{1, 2, 3}, []any{1, 2, 3}},
		{[]any{1, 2}, []any{1, nil, 2}},
		{[]any{1, 2, 1}, []any{1, 1, 2}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			p, q := CreateTree(c.P), CreateTree(c.Q)
			assert.Equal(t, slices.Equal(c.P, c.Q), isSameTree(p, q))
		})
	}
}

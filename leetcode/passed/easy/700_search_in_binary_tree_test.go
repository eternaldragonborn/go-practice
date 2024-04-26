package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	p := root
	for {
		if p == nil {
			return nil
		}
		if p.Val == val {
			return p
		}

		if val > p.Val {
			p = p.Right
		} else {
			p = p.Left
		}
	}
}

func TestSearchBST(t *testing.T) {
	cases := []struct {
		Data     []any
		Val      int
		Excepted []any
	}{
		{[]any{4, 2, 7, 1, 3}, 2, []any{2, 1, 3}},
		{[]any{4, 2, 7, 1, 3}, 5, []any{}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, searchBST(root, c.Val).PreOrderTraversal())
		})
	}
}

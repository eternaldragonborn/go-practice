package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return root
}

func TestInvertTree(t *testing.T) {
	cases := []struct {
		Data, Excepted []any
	}{
		{[]any{4, 2, 7, 1, 3, 6, 9}, []any{4, 7, 2, 9, 6, 3, 1}},
		{[]any{2, 1, 3}, []any{2, 3, 1}},
		{[]any{}, []any{}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, invertTree(root).PreOrderTraversal())
		})
	}
}

package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func goodNodes(root *TreeNode) int {
	goodNodes := 0
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}

		if node.Val >= max {
			goodNodes++
			max = node.Val
		}

		dfs(node.Right, max)
		dfs(node.Left, max)
	}

	dfs(root, math.MinInt64)

	return goodNodes
}

func TestGoodNodes(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted int
	}{
		{[]any{3, 1, 4, 3, nil, 1, 5}, 4},
		{[]any{3, 3, nil, 4, 2}, 3},
		{[]any{1}, 1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, goodNodes(root))
		})
	}
}

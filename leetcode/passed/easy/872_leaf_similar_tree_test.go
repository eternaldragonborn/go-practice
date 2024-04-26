package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := []int{}, []int{}

	var dfs func(node *TreeNode, leaves *[]int)
	dfs = func(node *TreeNode, leaves *[]int) {
		if node.Left == nil && node.Right == nil {
			// fmt.Println(node.Val)
			(*leaves) = append((*leaves), node.Val)
			return
		}

		if node.Left != nil {
			dfs(node.Left, leaves)
		}
		if node.Right != nil {
			dfs(node.Right, leaves)
		}
	}

	dfs(root1, &leaves1)
	dfs(root2, &leaves2)

	// fmt.Println(leaves1, leaves2)
	return slices.Equal(leaves1, leaves2)
}

func TestLeafSimilar(t *testing.T) {
	cases := []struct {
		Data1, Data2 []any
		Excepted     bool
	}{
		{[]any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}, []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}, true},
		{[]any{1, 2, 3}, []any{1, 3, 2}, false},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root1 := CreateTree(c.Data1)
			root2 := CreateTree(c.Data2)

			assert.Equal(t, c.Excepted, leafSimilar(root1, root2))
		})
	}
}

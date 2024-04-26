package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func dfs(node *TreeNode, path *[]rune, smallest *string) {
	if node == nil {
		return
	}
	*path = append(*path, rune(node.Val)+'a')

	// is leaf node
	if node.Left == nil && node.Right == nil {
		var pathstr strings.Builder
		for i := len(*path) - 1; i >= 0; i-- {
			pathstr.WriteRune((*path)[i])
		}
		if str := pathstr.String(); *smallest == "" || str < *smallest {
			*smallest = str
		}
	}

	dfs(node.Left, path, smallest)
	dfs(node.Right, path, smallest)

	// remove current's character(backtrack)
	*path = (*path)[:len(*path)-1]
}

func smallestFromLeaf(root *TreeNode) string {
	var smallest string
	var path []rune

	dfs(root, &path, &smallest)

	return smallest
}

func TestSmallestFromLeaf(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted string
	}{
		{[]any{0, 1, 2, 3, 4, 3, 4}, "dba"},
		{[]any{25, 1, 3, 1, 3, 0, 2}, "adz"},
		{[]any{2, 2, 1, nil, 1, 0, nil, 0}, "abc"},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)

			assert.Equal(t, c.Excepted, smallestFromLeaf(root))
		})
	}
}

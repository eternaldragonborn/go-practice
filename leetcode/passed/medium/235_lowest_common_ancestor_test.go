package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// it's b-search tree, left child nodes are always <= current node
	// right nodes are always >= current node
	cur := root

	for {
		fmt.Println(cur.Val)
		if cur.Val > p.Val && cur.Val > q.Val {
			// both p and q are less than current node, LCA is in left subtree
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			// both are > cur.val, LCA is in right subtree
			cur = cur.Right
		} else {
			// two nodes are in different side or one of them is LCA
			return cur
		}
	}
}

func TestLoestCommonAncestor(t *testing.T) {
	cases := []struct {
		Data     []any
		P, Q     int
		Excepted int
		Name     string
	}{
		{[]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 8, 6, "root"},
		{[]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 3, 9, 6, "far"},
		{[]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 3, 5, 4, "deep"},
		{[]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 4, 2, "one is parent"},
		{[]any{2, 1}, 2, 1, 2, "two nodes"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := CreateTree(c.Data)
			p, q, LCA := root.FindNode(c.P), root.FindNode(c.Q), root.FindNode(c.Excepted)
			assert.Equal(t, LCA, lowestCommonAncestor(root, p, q))
		})
	}
}

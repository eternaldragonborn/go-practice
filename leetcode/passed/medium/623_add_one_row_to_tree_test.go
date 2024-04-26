package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newNode := new(TreeNode)
		newNode.Val = val
		newNode.Left = root
		return newNode
	}
	depthCount := 1
	nodeQueue := []*TreeNode{root}

	for {
		if depthCount == depth-1 {
			for _, node := range nodeQueue {
				newLeftNode, newRightNode := new(TreeNode), new(TreeNode)
				newLeftNode.Val, newRightNode.Val = val, val
				newLeftNode.Left = node.Left
				newRightNode.Right = node.Right

				node.Left, node.Right = newLeftNode, newRightNode
			}
			break
		}

		// nodes := make([]*TreeNode, 0, int(math.Pow(2, float64(depth-1)))) // ! error when depth of tree > 64(num of nodes > 2^64)
		nodes := []*TreeNode{}
		for _, node := range nodeQueue {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		nodeQueue = nodes
		depthCount++
	}

	return root
}

func TestAddOneRow(t *testing.T) {
	cases := []struct {
		Data       []any
		Val, Depth int
		Excepted   []any
	}{
		{[]any{4, 2, 6, 3, 1, 5}, 1, 2, []any{4, 1, 1, 2, nil, nil, 6, 3, 1, 5}},
		{[]any{4, 2, nil, 3, 1}, 1, 3, []any{4, 2, nil, 1, 1, 3, nil, nil, 1}},
		{[]any{4, 2, nil, 3}, 1, 3, []any{4, 2, nil, 1, 1, 3}},
		{[]any{4, 2, 6, 3, 1, 5}, 1, 3, []any{4, 2, 6, 1, 1, 1, 1, 3, nil, nil, 1, 5}},
		{[]any{4, 2, 6, 3, 1, 5}, 1, 1, []any{1, 4, nil, 2, 6, 3, 1, 5}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			ans := addOneRow(root, c.Val, c.Depth)
			assert.Equal(t, c.Excepted, ans.PreOrderTraversal())
		})
	}
}

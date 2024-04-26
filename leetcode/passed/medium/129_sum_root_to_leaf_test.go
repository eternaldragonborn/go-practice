package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	numbers := []*TreeNode{}
	numbers = append(numbers, root)

	for len(numbers) != 0 {
		node := numbers[0]
		// numbers[0] = nil
		numbers = numbers[1:]
		if node.Left == nil && node.Right == nil {
			sum += node.Val
			// fmt.Println(node.Val)
			continue
		}
		if node.Left != nil {
			node.Left.Val += node.Val * 10
			numbers = append(numbers, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val * 10
			numbers = append(numbers, node.Right)
		}

	}

	return sum
}

func TestSumLeafNumbers(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted int
	}{
		{[]any{1, 2, 3}, 25},
		{[]any{4, 9, 0, 5, 1}, 1026},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, sumNumbers(root))
		})
	}
}

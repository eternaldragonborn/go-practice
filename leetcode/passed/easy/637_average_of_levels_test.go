package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	averages := []float64{}
	// depth := 0

	for len(queue) != 0 {
		newQueue := []*TreeNode{}
		sum := 0.
		// nodeCount := 0

		for _, node := range queue {
			sum += float64(node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		averages = append(averages, sum/float64(len(queue)))
		queue = newQueue
	}

	return averages
}

func TestAverageOfLevels(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted []float64
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, []float64{3.00000, 14.50000, 11.00000}},
		{[]any{3, 9, 20, 15, 7}, []float64{3.00000, 14.50000, 11.00000}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, averageOfLevels(root))
		})
	}
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "leetcode.com/fafnir/structure/tree"
)

type queueNode struct {
	node *TreeNode
	next *queueNode
}
type treeNodeQueue struct {
	head *queueNode
}

func (q *treeNodeQueue) Push(node *TreeNode) {
	newNode := new(queueNode)
	newNode.node = node

	if q.head == nil {
		q.head = newNode
		return
	}

	var p = q.head
	for p.next != nil {
		p = p.next
	}
	p.next = newNode
}

func (q *treeNodeQueue) Pop() *TreeNode {
	p := q.head

	q.head = p.next

	return p.node
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	queue := new(treeNodeQueue)

	queue.Push(root)

	for queue.head != nil {
		ele := queue.Pop()

		if ele.Left != nil {
			if ele.Left.Left == nil && ele.Left.Right == nil {
				sum += ele.Left.Val
			} else {
				queue.Push(ele.Left)
			}
		}

		if ele.Right != nil {
			queue.Push(ele.Right)
		}
	}

	return sum
}

func TestSumOfLeftNode(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted int
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, 24},
		{[]any{1}, 0},
		{[]any{}, 0},
		{[]any{1, 2, 3, 4, 5}, 4},
		{[]any{1, 2, 2, 3, 3, nil, nil, 4, 4}, 4},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, sumOfLeftLeaves(root))
		})
	}
}

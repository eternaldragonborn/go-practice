package tree

type queueNode struct {
	val  *TreeNode
	next *queueNode
}

type queue struct {
	head *queueNode
}

func newQueue() *queue {
	q := new(queue)
	return q
}

func (q *queue) Push(val *TreeNode) {
	newNode := new(queueNode)
	newNode.val = val
	if q.head == nil {
		q.head = newNode
		return
	}

	var p *queueNode
	for p = q.head; p.next != nil; p = p.next {
	}

	p.next = newNode
}

func (q *queue) Pop() *TreeNode {
	p := q.head

	q.head = p.next

	return p.val
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (root *TreeNode) FindNode(val int) *TreeNode {
	nodeQueue := []*TreeNode{root}

	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if node.Val == val {
			return node
		}

		if left := node.Left; left != nil {
			nodeQueue = append(nodeQueue, left)
		}
		if right := node.Right; right != nil {
			nodeQueue = append(nodeQueue, right)
		}
	}

	return nil
}

func (root *TreeNode) LevelOrderTraversal() []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	var queue = newQueue()
	queue.Push(root)

	for queue.head != nil {
		node := queue.Pop()
		ans = append(ans, node.Val)

		if left := node.Left; left != nil {
			queue.Push(left)
		}
		if right := node.Right; right != nil {
			queue.Push(right)
		}
	}

	return ans
}

func (root *TreeNode) PostOrderTraversal() []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	return ans
}

func (root *TreeNode) PreOrderTraversal() []any {
	ans := []any{}
	if root == nil {
		return ans
	}

	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		if node != nil {
			ans = append(ans, node.Val)
			if node.Left != nil || node.Right != nil {
				nodeQueue = append(nodeQueue, node.Left)
				nodeQueue = append(nodeQueue, node.Right)
			}
		} else if len(nodeQueue) != 1 {
			ans = append(ans, nil)
		}

		nodeQueue[0] = nil
		nodeQueue = nodeQueue[1:]
	}

	return ans
}

func CreateTree(data []any) *TreeNode {
	var root *TreeNode
	if len(data) == 0 {
		return root
	}

	var queue = newQueue()

	root = &TreeNode{Val: data[0].(int)}
	queue.Push(root)

	for i := 1; i < len(data); i++ {
		ele := queue.Pop()
		// left node
		if x, ok := data[i].(int); ok {
			// x is int
			leftNode := new(TreeNode)
			leftNode.Val = x
			ele.Left = leftNode

			queue.Push(leftNode)
		}

		// right node
		i++
		if i == len(data) {
			break
		}
		if x, ok := data[i].(int); ok {
			rightNode := new(TreeNode)
			rightNode.Val = x
			ele.Right = rightNode

			queue.Push(rightNode)
		}
	}

	return root
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(treeHeight(root.Left), treeHeight(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}

	return diff <= 1
}

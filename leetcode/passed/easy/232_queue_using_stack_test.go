package leetcode

type queueNode struct {
	val  int
	next *queueNode
}

type MyQueue struct {
	head *queueNode
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	node := new(queueNode)
	node.val = x

	if this.head == nil {
		this.head = node
		return
	}

	p := this.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
}

func (this *MyQueue) Pop() int {
	val := this.head.val
	this.head = this.head.next
	return val
}

func (this *MyQueue) Peek() int {
	return this.head.val
}

func (this *MyQueue) Empty() bool {
	return this.head == nil
}

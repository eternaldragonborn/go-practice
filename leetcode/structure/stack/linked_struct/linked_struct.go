package linked_struct

import (
	"errors"
	"sync"
)

type node struct {
	Val  interface{}
	Next *node
}

type Stack struct {
	top  *node
	Len  int
	Lock sync.Mutex
}

func NewStack() *Stack {
	stack := new(Stack)
	return stack
}

func (s *Stack) Push(item interface{}) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	// s.items = append(s.items, item)
	newNode := new(node)
	newNode.Val = item
	newNode.Next = s.top

	s.top = newNode
	s.Len++
}

func (s *Stack) Pop() (interface{}, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.top == nil {
		return nil, errors.New("stack is empty")
	}

	val := s.top.Val
	s.top = s.top.Next
	s.Len--

	return val, nil
}

func (s *Stack) Peek() (interface{}, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.top == nil {
		return nil, errors.New("stack is empty")
	}

	return s.top.Val, nil
}

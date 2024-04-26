package slice

import (
	"errors"
	"sync"
)

type Stack struct {
	items []interface{}
	lock  sync.Mutex
}

func NewStack() *Stack {
	stack := new(Stack)
	// stack.items = []interface{}{}
	return stack
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[0 : len(s.items)-1]
		return item, nil
	}
}

func (s *Stack) Peek() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

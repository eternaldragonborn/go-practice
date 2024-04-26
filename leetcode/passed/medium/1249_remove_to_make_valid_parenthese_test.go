package leetcode

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val  int
	Next *Node
}
type Stack struct {
	top *Node
	len int
}

func (s *Stack) Push(v int) {
	node := new(Node)
	node.Val = v

	// if s.top == nil {
	// 	s.top = node
	// } else {
	node.Next = s.top
	s.top = node
	// }
	s.len++
}
func (s *Stack) Pop() (int, error) {
	if s.len == 0 {
		return 0, errors.New("stack is empty")
	}

	val := s.top.Val
	// fmt.Println(s.top)
	s.top = s.top.Next
	s.len--

	return val, nil
}
func (s *Stack) Peek() (int, error) {
	if s.len == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.top.Val, nil
}

func minRemoveToMakeValid_structStack(s string) string {
	stack := Stack{}
	removePos := map[int]struct{}{}

	for index, c := range s {
		if c == '(' {
			stack.Push(index)
		} else if c == ')' {
			if stack.len == 0 {
				// invalid ')'
				// removePos = append(removePos, index)
				removePos[index] = struct{}{}
			} else {
				// paired
				stack.Pop()
			}
		}
	}

	// remove redundant '(' first
	for stack.len != 0 {
		// fmt.Println(pos)
		pos, _ := stack.Pop()
		// removePos = append(removePos, pos)
		removePos[pos] = struct{}{}
	}
	// fmt.Println(removePos)

	// result := ""
	result := make([]rune, 0, len(s))
	for index, c := range s {
		if _, exist := removePos[index]; !exist {
			// result += string(c)
			result = append(result, c)
		}
	}

	return string(result)
}

func minRemoveToMakeValid_sliceStack(s string) string {
	removePos := []int{}

	for index, c := range s {
		if c == '(' {
			removePos = append(removePos, index)
		} else if c == ')' {
			// invalid ')'
			if len(removePos) == 0 || s[removePos[len(removePos)-1]] != '(' {
				removePos = append(removePos, index)
			} else {
				// paired
				removePos = removePos[:len(removePos)-1]
			}
		}
	}

	// fmt.Println(removePos)

	// result := ""	// ! string manipulating cost much more than []rune in both time and space
	result := make([]rune, 0, len(s))
	i := 0
	for index, c := range s {
		if i < len(removePos) && index == removePos[i] {
			i++
		} else {
			// result += string(c)
			result = append(result, c)
		}
	}

	// return result
	return string(result)
}

var minRemoveCases = []struct {
	S, Excepted string
}{
	{"lee(t(c)o)de)", "lee(t(c)o)de"},
	{"a)b(c)d", "ab(c)d"},
	{"a(b(c)d", "ab(c)d"},
	{"))((", ""},
	{"))(a)(", "(a)"},
}

func TestStructStack(t *testing.T) {
	for _, c := range minRemoveCases {
		t.Run("", func(t *testing.T) {
			ans := minRemoveToMakeValid_structStack(c.S)
			assert.Equal(t, c.Excepted, ans, c.S)
		})
	}
}

func TestSliceStack(t *testing.T) {
	for _, c := range minRemoveCases {
		t.Run("", func(t *testing.T) {
			ans := minRemoveToMakeValid_sliceStack(c.S)
			assert.Equal(t, c.Excepted, ans, c.S)
		})
	}
}

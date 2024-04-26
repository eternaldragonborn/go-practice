package leetcode

import (
	"errors"
	"slices"
	"testing"
)

type RuneStack struct {
	items []rune
}

func NewRuneStack() *RuneStack {
	stack := new(RuneStack)
	stack.items = []rune{}
	return stack
}
func (s *RuneStack) Push(item rune) { s.items = append(s.items, item) }
func (s *RuneStack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[0 : len(s.items)-1]
		return item, nil
	}
}

func isValid(s string) bool {
	stack := NewRuneStack()
	mapping := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack.Push(c)
		} else {
			item, err := stack.Pop()
			if err != nil || item != mapping[c] {
				return false
			}
		}
	}

	return len(stack.items) == 0
}

func TestRuneStack(t *testing.T) {
	stack := NewRuneStack()

	pushCases := []struct {
		C        rune
		Expected []rune
	}{
		{'a', []rune{'a'}},
		{'b', []rune{'a', 'b'}},
	}

	t.Run("push", func(t *testing.T) {
		for _, c := range pushCases {
			stack.Push(c.C)
			if !slices.Equal(stack.items, c.Expected) {
				t.Fatalf("%c pushed, expected %v, got %v", c.C, c.Expected, stack.items)
			}
		}
	})

	t.Run("pop", func(t *testing.T) {
		expected := []struct {
			Rune rune
			Err  bool
			Len  int
		}{
			{'b', false, 1},
			{'a', false, 0},
			{0, true, 0},
		}

		for _, c := range expected {
			result, err := stack.Pop()
			if c.Err && err == nil {
				t.Fatal("should be error but got", result)
			}

			if c.Rune != result {
				t.Fatalf("should pop %c, but got %c", c.Rune, result)
			}

			if c.Len != len(stack.items) {
				t.Fatalf("len should be %d, got %d instead", c.Len, len(stack.items))
			}
		}
	})
}

func TestParenthesis_20(t *testing.T) {
	cases := []struct {
		S        string
		Expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"({)}", false},
		{"({}([]))", true},
	}

	for _, c := range cases {
		t.Run("valid", func(t *testing.T) {
			if ans := isValid(c.S); ans != c.Expected {
				t.Fatalf("%s expected %v, got %v", c.S, c.Expected, ans)
			}
		})
	}
}

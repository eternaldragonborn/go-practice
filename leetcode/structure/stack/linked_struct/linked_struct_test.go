package linked_struct

import (
	"testing"
)

func lenOfStack(s *Stack) int {
	len := 0
	for p := s.top; p != nil; p = p.Next {
		len++
	}

	return len
}

func TestLinkedStructStack(t *testing.T) {
	s := NewStack()
	cases := []struct {
		Name   string
		Action string
		Data   interface{}
		Err    bool
	}{
		{"empty pop", "pop", nil, true},
		{"empty peek", "peek", nil, true},
		{"push 1", "push", 1, false},
		{"push 2", "push", 2, false},
		{"pop 2", "pop", 2, false},
		{"peek 1", "peek", 1, false},
		{"pop 1", "pop", 1, false},
		{"empty", "pop", nil, true},
		{"push 3", "push", 3, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if c.Action == "pop" {
				ans, err := s.Pop()
				if c.Err && err == nil {
					t.Fatalf("error excepted")
				}

				if ans != c.Data {
					t.Fatalf("%v excepted, got %v", c.Data, ans)
				}
			} else if c.Action == "push" {
				s.Push(c.Data)
			} else if c.Action == "peek" {
				ans, err := s.Peek()
				if c.Err && err == nil {
					t.Fatalf("error excepted")
				}

				if ans != c.Data {
					t.Fatalf("%v excepted, got %v", c.Data, ans)
				}
			}

			if excepted := lenOfStack(s); s.Len != excepted {
				t.Fatalf("excepted len: %d(got %d)", excepted, s.Len)
			}
			t.Logf("stack len: %d", s.Len)
		})
	}
}

func BenchmarkPush(b *testing.B) {
	s := NewStack()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewStack()

	// push
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	// pop
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

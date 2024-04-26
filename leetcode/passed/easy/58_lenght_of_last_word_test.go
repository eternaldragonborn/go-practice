package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	last_word := words[len(words)-1]

	fmt.Println(last_word)

	return len(last_word)
}

func TestLastWordLength(t *testing.T) {
	cases := []struct {
		S      string
		Except int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	t.Run("test length of last word", func(t *testing.T) {
		for _, c := range cases {
			if ans := lengthOfLastWord(c.S); ans != c.Except {
				t.Fatalf("excepted %d, got %d", c.Except, ans)
			}
		}
	})
}

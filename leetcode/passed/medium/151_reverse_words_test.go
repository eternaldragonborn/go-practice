package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseWords(s string) string {
	i := 0
	words := []string{}

	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}

		var word strings.Builder
		for i < len(s) && s[i] != ' ' {
			word.WriteByte(s[i])
			i++
		}
		if word.Len() != 0 {
			words = append(words, word.String())
		}
	}

	var ans strings.Builder
	for i = len(words) - 1; i >= 0; i-- {
		ans.WriteString(words[i])
		if i != 0 {
			ans.WriteByte(' ')
		}
	}
	return ans.String()
}

func TestReverseWords(t *testing.T) {
	cases := []struct {
		S, Excepted string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, reverseWords(c.S))
		})
	}
}

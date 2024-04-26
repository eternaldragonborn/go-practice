package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordPattern(pattern string, s string) bool {
	type void struct{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	dict := map[rune]string{}
	wordDict := map[string]void{}

	for index, c := range pattern {
		word, exist := dict[c]
		_, wordExist := wordDict[words[index]]
		if (exist && word != words[index]) || (!exist && wordExist) {
			return false
		}

		if !exist {
			dict[c] = words[index]
			wordDict[words[index]] = void{}
		}
	}

	return true
}

func TestWordPattern(t *testing.T) {
	cases := []struct {
		Pattern, S string
		Excepted   bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog dog dog dog", false},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, wordPattern(c.Pattern, c.S))
		})
	}
}

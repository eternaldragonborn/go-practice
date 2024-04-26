package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := make(map[rune]int)

	for _, c := range s {
		dict[c]++
	}

	for _, c := range t {
		if _, exist := dict[c]; !exist {
			return false
		}
		dict[c]--
		if dict[c] == 0 {
			delete(dict, c)
		}
	}

	return len(dict) == 0
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		S, T     string
		Excepted bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "cat", false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, isAnagram(c.S, c.T))
		})
	}
}

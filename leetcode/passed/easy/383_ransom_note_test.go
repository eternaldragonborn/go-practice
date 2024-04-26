package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// O(n + m)
func canConstruct(ransomNote string, magazine string) bool {
	lookupLetters := map[rune]int{}
	for _, c := range magazine {
		lookupLetters[c]++
	}

	for _, c := range ransomNote {
		if t, exist := lookupLetters[c]; !exist || t == 0 {
			return false
		}

		lookupLetters[c]--
	}

	return true
}

func TestRansomNote(t *testing.T) {
	cases := []struct {
		Note, Magazine string
		Excepted       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, c.Excepted, canConstruct(c.Note, c.Magazine), "note: %s, magazine: %s", c.Note, c.Magazine)
		})
	}
}

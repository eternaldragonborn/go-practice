package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	i := 0
	for _, c := range t {
		if c == rune(s[i]) {
			i++
		}

		if i == len(s) {
			return true
		}
	}

	return false
}

func TestSubsequence(t *testing.T) {
	cases := []struct {
		S, T     string
		Excepted bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, isSubsequence(c.S, c.T))
		})
	}
}

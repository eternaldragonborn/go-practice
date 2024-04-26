package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isVowel(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c == 'a' ||
		c == 'e' ||
		c == 'i' ||
		c == 'o' ||
		c == 'u'
}

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	l, r := 0, len(s)-1
	var ans strings.Builder

	for l < len(s) && r >= 0 {
		for l < len(s) && !isVowel(s[l]) {
			ans.WriteByte(s[l])
			l++
		}
		for r > 0 && !isVowel(s[r]) {
			r--
		}

		if r >= 0 && isVowel(s[r]) {
			ans.WriteByte(s[r])
		}
		l++
		r--
	}

	for l < len(s) {
		ans.WriteByte(s[l])
		l++
	}

	return ans.String()
}

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		S, Excepted string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"ehalo", "ohale"},
		{"a.", "a."},
		{",.", ",."},
		{".a", ".a"},
		{"Aa", "aA"},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, reverseVowels(c.S))
		})
	}
}

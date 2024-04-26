package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		// find next alphabet
		for ; !isAlphanumeric(s[l]) && l < r; l++ {
		}
		for ; !isAlphanumeric(s[r]) && r > l; r-- {
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func TestValidPalindrome(t *testing.T) {
	cases := []struct {
		Name     string
		S        string
		Excepted bool
	}{
		{"is palindrome", "A man, a plan, a canal: Panama", true},
		{"not palindrome", "race a car", false},
		{"empty string", " ", true},
		{"no alphabet", ".,", true},
		{"", ".,a", true},
		{"", "., a., ", true},
		{"", "0P", false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, isPalindrome(c.S), c.S)
		})
	}
}

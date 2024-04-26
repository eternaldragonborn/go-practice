package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isDivisor(s, t string) bool {
	for i := 0; i < len(s); i += len(t) {
		if s[i:i+len(t)] != t {
			return false
		}
	}
	return true
}

func isDivisor2(s1, s2, t string) bool {
	for strings.HasPrefix(s1, t) {
		s1 = strings.TrimPrefix(s1, t)
	}
	if s1 != "" {
		return false
	}

	for strings.HasPrefix(s2, t) {
		s2 = strings.TrimPrefix(s2, t)
	}
	return s2 == ""
}

func gcdOfStrings(str1 string, str2 string) string {
	n := min(len(str1), len(str2))

	for n > 0 {
		// find gcd of both strings' length
		for len(str1)%n != 0 || len(str2)%n != 0 && n > 0 {
			n--
		}

		t := str1[:n]
		// if isDivisor(str1, t) && isDivisor(str2, t) {
		// 	return t
		// }
		if isDivisor2(str1, str2, t) {
			return t
		}
		n--
	}
	return ""
}

func TestGcdOfString(t *testing.T) {
	cases := []struct {
		Name, Str1, Str2, Excepted string
	}{
		{"no repeat", "ABCABC", "ABC", "ABC"},
		{"repeat", "ABABAB", "ABAB", "AB"},
		{"repeat", "ABAB", "ABABAB", "AB"},
		{"no common divisor", "LEET", "CODE", ""},
		{"repeat letter", "ABAABAABA", "ABAABA", "ABA"},
		{"repeat letter and repeat divisor", "ABAABAABAABA", "ABAABA", "ABAABA"},
		{"single letter of divisor", "AAA", "AA", "A"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, gcdOfStrings(c.Str1, c.Str2))
		})
	}
}

func TestIsDivisor(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(true, isDivisor("ABCABC", "ABC"))
	assertion.Equal(true, isDivisor("ABCABCABC", "ABC"))
	assertion.Equal(false, isDivisor("ABCABC", "AB"))
}

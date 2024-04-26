package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func isUpper(c byte) (bool, byte) {
	// upper case
	if c >= 'A' && c <= 'Z' {
		return true, c - 'A'
	} else { // lower case
		return false, c - 'a'
	}
}

func makeGood(s string) string {
	for index := 0; index < len(s)-1; index++ {
		isUpper1, alphabet1 := isUpper(s[index])
		isUpper2, alphabet2 := isUpper(s[index+1])

		// is vice-versa
		if isUpper1 != isUpper2 && alphabet1 == alphabet2 {
			s = strings.ReplaceAll(s, fmt.Sprintf("%c%c", s[index], s[index+1]), "")
			index = -1
		}
	}

	return s
}

func TestAns(t *testing.T) {
	cases := []struct {
		S, Excepted string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
		{"s", "s"},
	}

	for _, c := range cases {
		t.Run(c.S, func(t *testing.T) {
			if ans := makeGood(c.S); ans != c.Excepted {
				t.Errorf("excepted %s, got %s", c.Excepted, ans)
			}
		})
	}
}

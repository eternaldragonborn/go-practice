package leetcode

import (
	"fmt"
	"testing"
)

func isIsomorphic(s string, t string) bool {
	mapping1 := map[rune]rune{}
	mapping2 := map[rune]struct{}{}

	if len(s) != len(t) {
		return false
	}

	for i, c1 := range s {
		c2 := rune(t[i])

		mappedC2, exist1 := mapping1[c1]
		_, exist2 := mapping2[c2]

		if exist1 != exist2 || (exist1 && mappedC2 != c2) {
			return false
		}
		// first appear
		if !exist1 {
			mapping1[c1] = c2
			mapping2[c2] = struct{}{}
		}
	}

	return true
}

func TestIsomorphic(t *testing.T) {
	cases := []struct {
		S, T     string
		Excepted bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if ans := isIsomorphic(c.S, c.T); ans != c.Excepted {
				t.Log(c.S, c.T)
				t.Errorf("excepted %v, got %v", c.Excepted, ans)
			}
		})
	}
}

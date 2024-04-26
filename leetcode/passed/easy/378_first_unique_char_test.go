package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func firstUniqChar(s string) int {
	type void struct{}
	repeated := make(map[rune]void, len(s))
	nonRepeated := map[rune]int{}

	for index, c := range s {
		if _, exist := repeated[c]; exist {
			continue
		}

		if _, exist := nonRepeated[c]; exist {
			repeated[c] = void{}
			delete(nonRepeated, c)
		} else {
			nonRepeated[c] = index
		}
	}

	if len(nonRepeated) == 0 {
		return -1
	}

	ans := len(s)
	for _, v := range nonRepeated {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func TestFirstUniqueChar(t *testing.T) {
	cases := []struct {
		S        string
		Excepted int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}

	for _, c := range cases {
		t.Run(c.S, func(t *testing.T) {
			assert.Equal(t, c.Excepted, firstUniqChar(c.S))
		})
	}
}

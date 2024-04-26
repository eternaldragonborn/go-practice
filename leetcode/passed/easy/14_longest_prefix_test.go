package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	var ans strings.Builder

	i := 0
outer:
	for {
		flag := true
		if i >= len(strs[0]) {
			break
		}
		c := strs[0][i]

		for _, str := range strs[1:] {
			if i >= len(str) {
				break outer
			}
			if str[i] != c {
				flag = false
				break outer
			}
		}
		if flag {
			ans.WriteByte(c)
		}
		i++
	}
	fmt.Println(i)

	return ans.String()
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		Strs     []string
		Excepted string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, longestCommonPrefix(c.Strs))
		})
	}
}

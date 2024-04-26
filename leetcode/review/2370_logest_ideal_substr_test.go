package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)

	for _, c := range s {
		c := int(c - 'a')
		l := max(0, c-k)
		r := min(25, c+k)
		maxLen := 0

		for i := l; i <= r; i++ {
			maxLen = max(maxLen, dp[i])
		}
		dp[c] = maxLen + 1
	}

	// fmt.Println(dp)
	maxLen := 0
	for _, len := range dp {
		maxLen = max(maxLen, len)
	}
	return maxLen
}

func TestLongestIdealString(t *testing.T) {
	cases := []struct {
		S           string
		K, Excepted int
	}{
		{"acfgbd", 2, 4},
		{"abcd", 3, 4},
		{"eduktdb", 15, 5},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, longestIdealString(c.S, c.K))
		})
	}
}

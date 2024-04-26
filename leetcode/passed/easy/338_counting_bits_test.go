package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countBits(n int) []int {
	ans := make([]int, n+1)
	x := 1 // pow of 2

	for i := 1; i <= n; i++ {
		if i >= x*2 {
			x *= 2
		}
		ans[i] = ans[i-x] + 1
	}

	return ans
}

/*
1 1
2 2
3 2
4 4
5 4
6 4
7 4
8 8
*/

func TestCountBits(t *testing.T) {
	cases := []struct {
		N        int
		Excepted []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, countBits(c.N))
		})
	}
}

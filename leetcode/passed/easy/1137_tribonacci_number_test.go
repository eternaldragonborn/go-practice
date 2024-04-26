package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func tribonacci(n int) int {
	Tn := []int{0, 1, 1}
	if n < 3 {
		return Tn[n]
	}

	for i := 3; i <= n; i++ {
		temp := 0
		for _, n := range Tn {
			temp += n
		}
		copy(Tn, Tn[1:])
		Tn[2] = temp
	}
	return Tn[2]
}

func TestTribonacci(t *testing.T) {
	cases := []struct {
		Name        string
		N, Excepted int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"4", 4, 4},
		{"25", 25, 1389537},
	}

	t.Parallel()
	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, tribonacci(c.N))
		})
	}
}

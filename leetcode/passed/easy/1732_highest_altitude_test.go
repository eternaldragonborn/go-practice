package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestAltitude(gain []int) int {
	altitude, ans := 0, 0

	for _, x := range gain {
		altitude += x

		if altitude > ans {
			ans = altitude
		}
	}

	return ans
}

func TestLargestAltitude(t *testing.T) {
	cases := []struct {
		Gain     []int
		Excepted int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, largestAltitude(c.Gain))
		})
	}
}

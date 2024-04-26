package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	max := 0

	for _, n := range candies {
		if n > max {
			max = n
		}
	}

	for index, n := range candies {
		ans[index] = n+extraCandies >= max
	}

	return ans
}

func TestKidsWithCandies(t *testing.T) {
	cases := []struct {
		Candies      []int
		ExtraCandies int
		Excepted     []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, kidsWithCandies(c.Candies, c.ExtraCandies))
		})
	}
}

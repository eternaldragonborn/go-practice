package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ideally solution: https://ithelp.ithome.com.tw/articles/10213285
func majorityElement(nums []int) int {
	counting := map[int]int{}

	for _, n := range nums {
		counting[n]++

		if counting[n] > len(nums)/2 {
			return n
		}
	}

	return -1
}

func TestMajorityElement_169(t *testing.T) {
	cases := []struct {
		Nums     []int
		Excepted int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, majorityElement(c.Nums))
		})
	}
}

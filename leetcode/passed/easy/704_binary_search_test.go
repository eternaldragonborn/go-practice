package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func TestBSearch_704(t *testing.T) {
	cases := []struct {
		Nums             []int
		Target, Excepted int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, search(c.Nums, c.Target))
		})
	}
}

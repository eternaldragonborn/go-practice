package leetcode

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func bSearch(nums []int, start, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return start + mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

// 20.11 / 28.30
func search(nums []int, target int) int {
	if nums[0] > nums[len(nums)-1] { // rotated
		l, r := 0, len(nums)-1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] == target {
				return mid
			}

			if nums[mid] > nums[mid+1] { // [mid+1] is the start of rotated slice(minimum)
				if target < nums[0] {
					// target in rotated slice
					l, r = mid+1, len(nums)-1
				} else {
					// target in left part of slice
					l, r = 0, mid
				}

				return bSearch(nums[l:r+1], l, target)
			} else {
				if nums[mid] < nums[0] { // start of rotated slice is in left part
					r = mid
				} else {
					l = mid + 1
				}
			}
		}
	} else {
		return bSearch(nums, 0, target)
	}

	return -1
}

func TestRotatedArraySearch(t *testing.T) {
	cases := []struct {
		Nums             []int
		Target, Excepted int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{6, 7, 0, 1, 2, 3, 4}, 3, 5},
		{[]int{4, 5, 6, 7, 8, 1, 2}, 2, 6},
		{[]int{4, 5, 6, 7, 8, 1, 2}, 4, 0},
		{[]int{1, 2, 4, 6, 7, 11, 13}, 11, 5},
		{[]int{1, 2, 4, 6, 7, 11, 13}, 10, -1},
		{[]int{2, 1}, 1, 1},
		{[]int{2, 0, 1}, 1, 2},
		{[]int{2, 3, 1}, 1, 2},
		{[]int{1}, 0, -1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, search(c.Nums, c.Target), c.Nums)
		})
	}
}

func TestBSearch(t *testing.T) {
	n := rand.Intn(500) + 100
	nums := make([]int, n)

	for j := 0; j < n; j++ {
		nums[j] = j
	}

	cases := []struct {
		name   string
		target int
		found  bool
	}{
		{"first", 0, true},
		{"last", n - 1, true},
		{"random1", rand.Intn(n), true},
		{"random2", rand.Intn(n), true},
		{"random3", rand.Intn(n), true},
		{"random4", rand.Intn(n), true},
		{"random5", rand.Intn(n), true},
		{"not found", n + 100, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.found {
				assert.Equal(t, c.target, bSearch(nums, 0, c.target))
			} else {
				assert.Equal(t, -1, bSearch(nums, 0, c.target))
			}
		})
	}
}

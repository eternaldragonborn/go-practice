package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeElement(nums []int, val int) int {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
	}

	for i := l; i < len(nums); i++ {
		nums[i] = 0
	}
	// nums = nums[:l]

	return l
}

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		Nums, ExceptedNums []int
		Val, Excepted      int
	}{
		{[]int{3, 2, 2, 3}, []int{2, 2, 0, 0}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, []int{0, 1, 4, 0, 3, 0, 0, 0}, 2, 5},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ans := removeElement(c.Nums, c.Val)
			f1 := assert.Equal(t, c.Excepted, ans)
			f2 := assert.ElementsMatch(t, c.ExceptedNums, c.Nums)
			// f2 := assert.Equal(t, c.ExceptedNums, c.Nums)

			if f1 || f2 {
				t.Log("val: ", c.Val)
			}
		})
	}
}

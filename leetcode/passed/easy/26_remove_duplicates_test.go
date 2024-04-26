package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	// type void struct{}
	// lookupMap := map[int]void{}
	l := 0 // next position that is not duplicate number

	/* for _, x := range nums[1:] {
		if _, exist := lookupMap[x]; !exist {
			nums[l] = x
			l++
			lookupMap[x] = void{}
		}
	} */
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l] {
			l++
			nums[l] = nums[i]
		}
	}

	for i := l + 1; i < len(nums); i++ {
		nums[i] = 0
	}

	return l + 1
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		Nums, ExceptedNums []int
		Excepted           int
	}{
		{[]int{1, 1, 2}, []int{1, 2, 0}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}, 5},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{1}, []int{1}, 1},
		{[]int{1, 1}, []int{1, 0}, 1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ans := removeDuplicates(c.Nums)
			assert.Equal(t, c.Excepted, ans)
			assert.ElementsMatch(t, c.ExceptedNums, c.Nums)
		})
	}
}

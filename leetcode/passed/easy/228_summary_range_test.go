package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printRange(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	} else {
		return fmt.Sprintf("%d->%d", start, end)
	}
}

func summaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return ans
	}

	start, end := nums[0], nums[0]

	for _, n := range nums[1:] {
		if n == end+1 {
			end = n
		} else {
			ans = append(ans, printRange(start, end))

			start, end = n, n
		}
	}

	ans = append(ans, printRange(start, end))

	return ans
}

func TestSummaryRanges(t *testing.T) {
	cases := []struct {
		Nums     []int
		Excepted []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{0, 2, 3, 4, 6, 8, 9, 11}, []string{"0", "2->4", "6", "8->9", "11"}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ans := summaryRanges(c.Nums)
			assert.Equal(t, c.Excepted, ans)
		})
	}
}

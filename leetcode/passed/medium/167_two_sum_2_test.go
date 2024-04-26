package leetcode

import (
	"slices"
	"testing"
)

func twoSum(numbers []int, target int) []int {
	mapping := make(map[int]int)

	for index, x := range numbers {
		if index2, exist := mapping[x]; exist {
			return []int{index2, index + 1}
		}

		mapping[target-x] = index + 1
	}

	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Numbers  []int
		Target   int
		Excepted []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	t.Run("test two sum", func(t *testing.T) {
		for _, c := range cases {
			if ans := twoSum(c.Numbers, c.Target); !slices.Equal(ans, c.Excepted) {
				t.Fatalf("excepted %v, got %v", c.Excepted, ans)
			}
		}
	})
}

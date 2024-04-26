package leetcode

import (
	"slices"
	"testing"
)

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	result := [][]int{}
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0

	for i < len1 && j < len2 {
		n1, n2 := nums1[i], nums2[j]

		if n1[0] < n2[0] {
			result = append(result, n1)
			i++
		} else if n1[0] == n2[0] {
			result = append(result, []int{n1[0], n1[1] + n2[1]})
			i++
			j++
		} else {
			result = append(result, n2)
			j++
		}
	}

	for ; i < len1; i++ {
		result = append(result, nums1[i])
	}

	for ; j < len2; j++ {
		result = append(result, nums2[j])
	}

	return result
}

func TestMerge(t *testing.T) {
	cases := []struct {
		Nums1, Nums2, Excepted [][]int
	}{
		{[][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}, [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}}},
		{[][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}, [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := mergeArrays(c.Nums1, c.Nums2); !slices.EqualFunc(ans, c.Excepted, func(i1, i2 []int) bool {
				return slices.Equal(i1, i2)
			}) {
				t.Errorf("excepted %v, got %v", c.Excepted, ans)
			}
		})
	}
}

package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findDifference_table(nums1 []int, nums2 []int) [][]int {
	type void struct{}
	ans0, ans1 := make([]int, 0, len(nums2)), make([]int, 0, len(nums1))
	numTable1, numTable2 := map[int]void{}, map[int]void{}
	for _, x := range nums1 {
		numTable1[x] = void{}
	}

	for _, x := range nums2 {
		if _, exist2 := numTable2[x]; !exist2 {
			// not duplicate
			numTable2[x] = void{}

			// not found in nums1
			if _, exist1 := numTable1[x]; !exist1 {
				ans1 = append(ans1, x)
			}
		}

	}

	for x := range numTable1 {
		if _, exist := numTable2[x]; !exist {
			// not found in nums2
			ans0 = append(ans0, x)
		}
	}

	return [][]int{ans0, ans1}
}

func findDifference_contains(nums1 []int, nums2 []int) [][]int {
	ans0, ans1 := make([]int, 0, len(nums2)), make([]int, 0, len(nums1))
	for _, x := range nums1 {
		if !slices.Contains(nums2, x) {
			ans0 = append(ans0, x)
		}
	}

	for _, x := range nums2 {
		if !slices.Contains(nums1, x) {
			ans1 = append(ans1, x)
		}
	}

	return [][]int{ans0, ans1}
}

func testFindDifference(t *testing.T, f func([]int, []int) [][]int) {
	t.Helper()
	cases := []struct {
		Nums1, Nums2 []int
		Excepted     [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
		{
			[]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10},
			[]int{-1, -40, -44, 41, 10, -43, 69, 10, 2},
			[][]int{{-81, -35, -10, -28, -61, -45, -15, 14, -80, 63}, {-1, 2, 69, -40, 41, 10, -43, -44}},
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ans := f(c.Nums1, c.Nums2)

			assertion := assert.New(t)
			assertion.ElementsMatch(c.Excepted[0], ans[0], "ans[0] mismatch")
			assertion.ElementsMatch(c.Excepted[1], ans[1], "ans[1] mismatch")
		})
	}
}

func TestHashTable(t *testing.T) {
	testFindDifference(t, findDifference_table)
}

func TestSlicesContain(t *testing.T) {
	testFindDifference(t, findDifference_contains)
}

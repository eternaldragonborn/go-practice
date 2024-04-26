package leetcode

import (
	"slices"
	"testing"
)

func Intersection(nums1 []int, nums2 []int) []int {
	type void struct{}
	var ans = make([]int, 0, len(nums1))
	var numbers = make(map[int]void)

	for _, x := range nums1 {
		numbers[x] = void{}
	}

	for _, x := range nums2 {
		if _, exist := numbers[x]; exist {
			ans = append(ans, x)
			delete(numbers, x)
		}
	}

	return ans
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		Nums1, Nums2, Expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := Intersection(c.Nums1, c.Nums2); !slices.Equal(ans, c.Expected) {
				t.Fatalf("%v and %v expected %v, but got %v", c.Nums1, c.Nums2, c.Expected, ans)
			}
		})
	}
}

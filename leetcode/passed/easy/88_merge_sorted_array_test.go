package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1 // position the number put
	m, n = m-1, n-1

	// from bigger num to smaller
	for ; m >= 0 && n >= 0; idx-- {
		if nums1[m] > nums2[n] {
			nums1[idx] = nums1[m]
			m--
		} else {
			nums1[idx] = nums2[n]
			n--
		}
	}

	for ; m >= 0; m, idx = m-1, idx-1 {
		nums1[idx] = nums1[m]
	}

	for ; n >= 0; n, idx = n-1, idx-1 {
		nums1[idx] = nums2[n]
	}
}

func TestMergeArray_88(t *testing.T) {
	cases := []struct {
		Name                   string
		Nums1, Nums2, Excepted []int
		M, N                   int
	}{
		{
			Name:  "same length",
			Nums1: []int{1, 2, 3, 0, 0, 0},
			Nums2: []int{2, 5, 6},
			M:     3, N: 3,
			Excepted: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Name:  "empty num2",
			Nums1: []int{1},
			Nums2: []int{},
			M:     1, N: 0,
			Excepted: []int{1},
		},
		{
			Name:  "empty num1",
			Nums1: []int{0},
			Nums2: []int{1},
			M:     0, N: 1,
			Excepted: []int{1},
		},
		{
			Name:  "different length",
			Nums1: []int{1, 2, 3, 8, 9, 11, 13, 0, 0, 0, 0},
			Nums2: []int{2, 4, 9, 10},
			M:     7, N: 4,
			Excepted: []int{1, 2, 2, 3, 4, 8, 9, 9, 10, 11, 13},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			merge(c.Nums1, c.M, c.Nums2, c.N)
			assert.Equal(t, c.Excepted, c.Nums1)
		})
	}
}

package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt
	// sums := make([]int, len(nums)-k+1)
	sum := 0
	for index := range nums[:len(nums)-k+1] {
		if index == 0 {
			for _, n := range nums[:k] {
				// sums[0] += n
				sum += n
			}
		} else {
			// sums[index] = sums[index-1] - nums[index-1] + nums[index+k-1]
			sum = sum - nums[index-1] + nums[index+k-1]
		}
		// maxSum = max(sums[index], maxSum)
		if sum > maxSum {
			maxSum = sum
		}
	}
	// fmt.Println(sums)

	return float64(maxSum) / float64(k)
}

func TestMaxAverage_1(t *testing.T) {
	cases := []struct {
		Name     string
		Nums     []int
		K        int
		Excepted float64
	}{
		{"long", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"short", []int{1, 12, -5, -6, 50, 3}, 2, 26.5},
		{"k=len(nums)", []int{1, 12, -5, -6, 50, 3}, 6, 9.166666666},
		{"k=1", []int{1, 12, -5, -6, 50, 3}, 1, 50.},
		{"single element", []int{5}, 1, 5.},
		{"single negative", []int{-1}, 1, -1.},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			ans := findMaxAverage(c.Nums, c.K)
			assert.Lessf(t, math.Abs(ans-c.Excepted), 0.00001, "nums: %v, k= %d", c.Nums, c.K)
		})
	}
}

package leetcode

import (
	"math/rand"
	"testing"
)

func countQuadruplets_mapping(nums []int) int {
	// slices.Sort(nums) // ! wrong index order
	// fmt.Println(nums)
	count := 0
	mapping := map[int][]int{}

	for index, x := range nums {
		mapping[x] = append(mapping[x], index)
	}
	// fmt.Println(mapping)

	// !O(n^3)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				sum := nums[i] + nums[j] + nums[k]
				// if n, exist := mapping[sum]; exist {
				// 	fmt.Printf("%v, sum=%d, n=%d\n", []int{i, j, k}, sum, n)
				// 	count += mapping[sum]
				// }
				for _, l := range mapping[sum] {
					if l > k {
						count++
					}
				}
			}
		}
	}
	return count
}

func countQuadruplets_brute(nums []int) int {
	count := 0

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						count++
					}
				}
			}
		}
	}

	return count
}

var quadrupletsCases = []struct {
	Nums     []int
	Excepted int
}{
	{[]int{1, 2, 3, 6}, 1},
	{[]int{3, 3, 6, 4, 5}, 0},
	{[]int{1, 1, 1, 3, 5}, 4},
	{[]int{9, 6, 8, 23, 39, 23}, 2},
	{[]int{28, 8, 49, 85, 37, 90, 20, 8}, 1},
	{[]int{25, 69, 18, 24, 60, 7, 49}, 1},
}

func TestMapping(t *testing.T) {
	for _, c := range quadrupletsCases {
		t.Run("", func(t *testing.T) {
			t.Log(c.Nums)
			if ans := countQuadruplets_mapping(c.Nums); ans != c.Excepted {
				t.Errorf("excepted %d, got %d", c.Excepted, ans)
			}
		})
	}
}

func TestBrute(t *testing.T) {
	for _, c := range quadrupletsCases {
		t.Run("", func(t *testing.T) {
			c := c
			t.Parallel()
			t.Log(c.Nums)
			if ans := countQuadruplets_brute(c.Nums); ans != c.Excepted {
				t.Errorf("excepted %d, got %d", c.Excepted, ans)
			}
		})
	}
}

func generateCase(n int) []int {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}

	return nums
}

var n = 500

func BenchmarkQuadruplets(b *testing.B) {
	nums := generateCase(500)

	b.ResetTimer()
	b.Run("mapping", func(b *testing.B) {
		countQuadruplets_mapping(nums)
	})

	b.ResetTimer()
	b.Run("brute force", func(b *testing.B) {
		countQuadruplets_brute(nums)
	})
}

func BenchmarkMapping(b *testing.B) {
	nums := generateCase(n)
	b.ResetTimer()

	countQuadruplets_mapping(nums)
}

func BenchmarkBrute(b *testing.B) {
	nums := generateCase(n)
	b.ResetTimer()

	countQuadruplets_brute(nums)
}

package leetcode

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// O(n)
func rotate_slice(nums []int, k int) {
	k %= len(nums)
	result := make([]int, 0, len(nums))
	result = append(result, nums[len(nums)-k:]...)
	result = append(result, nums[:len(nums)-k]...)

	copy(nums, result)
}

// space complexity is O(1)
// time: O(n * k)
// ! timeout when array and k are very large
func rotate_for(nums []int, k int) {
	k %= len(nums)

	if len(nums)-k < k { // if move right is faster
		for i := 0; i < len(nums)-k; i++ {
			var temp int = nums[len(nums)-1]
			for index := len(nums) - 2; index >= 0; index-- {
				nums[index], temp = temp, nums[index]
			}
			nums[len(nums)-1] = temp
		}
	} else { // move right
		for i := 0; i < k; i++ {
			var temp int = nums[0]
			for index, x := range nums[1:] {
				nums[index+1], temp = temp, x
			}
			nums[0] = temp
		}
	}
}

// O(n)
func rotate_for2(nums []int, k int) {
	k %= len(nums)
	result := make([]int, len(nums))

	for index, x := range nums {
		idx := (index + k) % len(nums)
		result[idx] = x
	}

	copy(nums, result)
}

var cases_189 = []struct {
	Nums, Excepted []int
	K              int
}{
	{[]int{-1}, []int{-1}, 2},
	{[]int{1, 2, 3, 4, 5, 6, 7}, []int{5, 6, 7, 1, 2, 3, 4}, 3},
	{[]int{1, 2, 3, 4, 5, 6, 7}, []int{4, 5, 6, 7, 1, 2, 3}, 4},
	{[]int{-1, -100, 3, 99}, []int{3, 99, -1, -100}, 2},
}

func testRotate(t *testing.T, f func([]int, int)) {
	t.Helper()
	for _, c := range cases_189 {
		t.Run("", func(t *testing.T) {
			f(c.Nums, c.K)
			assert.Equal(t, c.Excepted, c.Nums)
		})
	}
}

func TestSliceConcat(t *testing.T) {
	testRotate(t, rotate_slice)
}

func TestForLoop(t *testing.T) {
	testRotate(t, rotate_for)
}

func TestForLoop_2(t *testing.T) {
	testRotate(t, rotate_for2)
}

func newArray(n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = i
	}

	return res
}

func BenchmarkArrayRotate(b *testing.B) {
	nums := newArray(100000)
	k := rand.Intn(100000)

	b.Run("for loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rotate_for(nums, rand.Intn(k))
		}
	})

	b.Run("for loop 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rotate_for2(nums, rand.Intn(k))
		}
	})

	b.Run("slice copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rotate_slice(nums, rand.Intn(k))
		}
	})
}

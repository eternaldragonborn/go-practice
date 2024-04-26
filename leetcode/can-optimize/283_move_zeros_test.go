package leetcode

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 27.26% / 11.10%
func sliceCopy(nums []int) {
	zeroCount := 0

	for index := 0; index < len(nums); index++ {
		if nums[index] == 0 {
			// idx := nums[len(nums)-1-zeroCount]
			copy(nums[index:], nums[index+1:])
			nums[len(nums)-1] = 0
			zeroCount++
			index--
		}

		if index == len(nums)-zeroCount-1 {
			break
		}
	}
}

// 16.40 / 14.20
func twoPointer(nums []int) {
	// next position that non-zero value is moving to
	l := 0

	for _, v := range nums {
		if v != 0 {
			nums[l] = v
			l++
		}
	}

	// have no non-zero value, make rest of position 0
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
}

func testMoveZeroes(t *testing.T, f func([]int)) {
	t.Helper()
	cases := []struct {
		Name           string
		Nums, Excepted []int
	}{
		{"zeroes", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"longer", []int{1, 2, 3, 54, 0, 324, 9, 0, 0, 0, 0, 23}, []int{1, 2, 3, 54, 324, 9, 23, 0, 0, 0, 0, 0}},
		{"longer ending with 0", []int{1, 2, 3, 54, 0, 324, 9, 0, 0, 0, 0, 23, 0, 0, 0, 0, 0}, []int{1, 2, 3, 54, 324, 9, 23, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"only zero", []int{0}, []int{0}},
		{"no zero", []int{1, 2, 5, 23, 5}, []int{1, 2, 5, 23, 5}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			f(c.Nums)
			assert.Equal(t, c.Excepted, c.Nums)
		})
	}
}

func TestSliceCopy(t *testing.T) {
	testMoveZeroes(t, sliceCopy)
}

func TestTwoPointer(t *testing.T) {
	testMoveZeroes(t, twoPointer)
}

func BenchmarkMoveZeros(b *testing.B) {
	newNums := func(n int) []int {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = rand.Intn(50)
		}
		return res
	}

	benchHelper := func(b *testing.B, name string, f func([]int), n int) {
		b.Helper()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// b.StopTimer()
				nums := newNums(n)
				// b.StartTimer()
				f(nums)
			}
		})
	}

	for _, n := range []int{1000, 10000, 100000, 500000} {
		benchHelper(b, "slice copy "+strconv.Itoa(n), sliceCopy, n)
		benchHelper(b, "two pointer "+strconv.Itoa(n), twoPointer, n)
	}
}

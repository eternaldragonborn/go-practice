package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minOperations(nums []int, k int) int {
	n := 0

	for _, x := range nums {
		n = x ^ n
	}
	n = n ^ k

	move := 0
	for i := 1; i <= n; i = i << 1 {
		// fmt.Printf("%b, %b, %b\n", i, n, i&n)
		if i&n != 0 {
			move++
		}
	}

	return move
}

func TestMinOperation(t *testing.T) {
	cases := []struct {
		Nums        []int
		K, Excepted int
	}{
		{[]int{2, 1, 3, 4}, 1, 2},
		{[]int{2, 0, 2, 0}, 0, 0},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, minOperations(c.Nums, c.K))
		})
	}
}

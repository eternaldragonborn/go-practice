package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isHappy(n int) bool {
	type void struct{}
	dict := map[int]void{}

	for {
		sum := 0

		for n != 0 {
			x := n % 10
			sum += (x * x)
			n /= 10
		}

		if sum == 1 {
			return true
		} else if _, exist := dict[sum]; exist {
			return false
		} else {
			dict[sum] = void{}
			n = sum
		}
	}
}

func TestIsHappy(t *testing.T) {
	cases := []struct {
		N        int
		Excepted bool
	}{
		{19, true},
		{2, false},
		{5468123, false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.N), func(t *testing.T) {
			assert.Equal(t, c.Excepted, isHappy(c.N))
		})
	}
}

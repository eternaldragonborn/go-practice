package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLucky(s string, k int) int {
	var sum, n int

	for _, c := range s {
		n = int(c - 'a' + 1)
		for n != 0 {
			sum += (n % 10)
			n /= 10
		}
	}

	for i := 1; i < k; i++ {
		sum, n = 0, sum
		for n != 0 {
			sum += (n % 10)
			n /= 10
		}
	}

	return sum
}

func TestSumOfString(t *testing.T) {
	cases := []struct {
		S        string
		K        int
		Excepted int
	}{
		{"zbax", 2, 8},
		{"leetcode", 2, 6},
		{"iiii", 1, 36},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, c.Excepted, getLucky(c.S, c.K), "s=%s, k=%d", c.S, c.K)
		})
	}
}

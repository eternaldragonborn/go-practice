package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}

	digits := []rune{}
	for _, n := range num {
		for len(digits) > 0 && n < digits[len(digits)-1] && k > 0 {
			digits = digits[:len(digits)-1]
			k--
		}

		if len(digits) > 0 || n != '0' {
			digits = append(digits, n)
		}
	}

	for k > 0 && len(digits) > 0 {
		digits = digits[:len(digits)-1]
		k--
	}

	if len(digits) == 0 {
		return "0"
	}

	return string(digits)
}

func TestRemoveDigits_402(t *testing.T) {
	cases := []struct {
		Num      string
		K        int
		Excepted string
	}{
		{"1342219", 1, "132219"},
		{"1432219", 3, "1219"},
		{"1432219", 4, "119"},
		{"1432219", 5, "11"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, c.Excepted, removeKdigits(c.Num, c.K), "num= %s, K= %d", c.Num, c.K)
		})
	}
}

package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome_array(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	digits := []int8{}
	for x != 0 {
		digits = append(digits, int8(x%10))
		x /= 10
	}

	for l, r := 0, len(digits)-1; l <= r; l, r = l+1, r-1 {
		if digits[l] != digits[r] {
			return false
		}
	}

	return true
}

func isPalindrome_array_presized(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	// digits := []int8{}
	digits := make([]int8, 0, int(math.Log10(float64(x)))+1)
	for x != 0 {
		digits = append(digits, int8(x%10))
		x /= 10
	}

	for l, r := 0, len(digits)-1; l <= r; l, r = l+1, r-1 {
		if digits[l] != digits[r] {
			return false
		}
	}

	return true
}

// ! wrong ans when there's 0 in the number(ex. 10021)
func isPalindrome_log(x int) bool {
	if x < 0 {
		return false
	}
	n := int(math.Log10(float64(x)))

	for {
		if x < 10 {
			return true
		}

		div := int(math.Pow10(n))
		l := x / div
		r := x % 10
		if l != r {
			return false
		}

		x %= div
		x /= 10
		n -= 2
	}
}

func testIsPalindrome(t *testing.T, f func(int) bool) {
	t.Helper()
	cases := []struct {
		Name     string
		X        int
		Excepted bool
	}{
		{"is palindrome(odd digits)", 121, true},
		{"is palindrome(even digits)", 9559, true},
		{"zero", 0, true},
		{"negative", -121, false},
		{"not palindrome(even digits)", 10, false},
		{"not palindrome(odd digits)", 951, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, f(c.X))
		})
	}
}

func TestDigitArray(t *testing.T) {
	testIsPalindrome(t, isPalindrome_array)
}

func TestLog(t *testing.T) {
	testIsPalindrome(t, isPalindrome_log)
}

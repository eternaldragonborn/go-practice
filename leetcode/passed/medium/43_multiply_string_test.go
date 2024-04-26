package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func multiply_lowToHigh(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	} else if num1 == "1" {
		return num2
	} else if num2 == "1" {
		return num1
	}

	n1, n2 := make([]int, len(num1)), make([]int, len(num2))
	ans := make([]int, max(len(num1), len(num2)), len(num1)+len(num2))

	for i, c := range num1 {
		n1[len(num1)-i-1] = int(c - '0')
	}
	for i, c := range num2 {
		n2[len(num2)-i-1] = int(c - '0')
	}

	// from lower digits to higher
	for i, x := range n1 {
		for j, y := range n2 {
			if i+j > len(ans)-1 {
				ans = append(ans, x*y)
			} else {
				ans[i+j] += x * y
			}
		}
	}

	var result string
	overflow := false
	for i, x := range ans {
		if x >= 10 {
			if i < len(ans)-1 {
				ans[i+1] += x / 10
			} else {
				overflow = true
				ans = append(ans, x/10)
			}

			ans[i] %= 10
		}
		result = fmt.Sprintf("%d%s", ans[i], result)
	}
	// fmt.Println(ans)
	if overflow {
		result = fmt.Sprintf("%d%s", ans[len(ans)-1], result)
	}

	return result
}

func multiply_highToLow(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	} else if num1 == "1" {
		return num2
	} else if num2 == "1" {
		return num1
	}

	ans := make([]int, len(num1)+len(num2))

	// start from low digit
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			ans[i+j+1] += mul
			ans[i+j] += (ans[i+j+1] / 10)
			ans[i+j+1] %= 10
		}
	}

	i := 0
	for ; ans[i] == 0; i++ {
	}
	var result strings.Builder
	ans = ans[i:]
	result.Grow(len(ans))
	for _, x := range ans {
		result.WriteByte(byte(x) + '0')
	}

	return result.String()
}

type TestCase_43 struct {
	Num1, Num2, Excepted string
}

func TestMultiply(t *testing.T) {
	cases := []TestCase_43{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"12", "345", "4140"},
		{"154678", "546849", "84585509622"},
		{"2845321594", "39877532156489726", "113464403360289604626943244"},
		{"2845321594894897894263215", "39877532156489726", "113464403395975924182199542856607487229090"},
		{"999", "9999", "9989001"},
		{"0", "16548912346498", "0"},
		{"16548912346498", "0", "0"},
		{"16548912346498", "1", "16548912346498"},
		{"1", "16548912346498", "16548912346498"},
	}

	testHelper := func(t *testing.T, f func(nums1, nums2 string) string) {
		t.Helper()
		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				assert.Equalf(t, c.Excepted, f(c.Num1, c.Num2), "%s, %s", c.Num1, c.Num2)
			})
		}
	}

	t.Run("low digit to high", func(t *testing.T) {
		testHelper(t, multiply_lowToHigh)
	})
	t.Run("high digit to low", func(t *testing.T) {
		testHelper(t, multiply_highToLow)
	})
}

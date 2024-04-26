package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	startIdx, length := 0, 0
	mapping := map[rune]int{}

	for index, c := range s {
		// the char have already appeared and it's index is after start of substring(in the substring)
		if prevIdx, exist := mapping[c]; exist && startIdx <= prevIdx {
			if length > maxLength {
				maxLength = length
			}
			// move the start index forward
			startIdx = prevIdx + 1
			// update the index in substring
			mapping[c] = index
			length = index - startIdx + 1
		} else {
			mapping[c] = index
			length++
		}
	}

	if length > maxLength {
		maxLength = length
	}

	return maxLength
}

func TestLongest(t *testing.T) {
	cases := []struct {
		S        string
		Excepted int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"zzxcvbasdfga", 10},
		{"dvdf", 3},
	}

	for _, c := range cases {
		t.Run(c.S, func(t *testing.T) {
			ans := lengthOfLongestSubstring(c.S)
			assert.Equal(t, c.Excepted, ans)
		})
	}
}

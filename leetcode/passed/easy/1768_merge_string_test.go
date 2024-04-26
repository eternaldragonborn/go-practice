package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	len1, len2 := len(word1), len(word2)
	var result strings.Builder
	result.Grow(len1 + len2)

	for i < len1 && j < len2 {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}

	for ; i < len1; i++ {
		result.WriteByte(word1[i])
	}
	for ; j < len2; j++ {
		result.WriteByte(word2[j])
	}

	return result.String()
}

func TestMergeString(t *testing.T) {
	cases := []struct {
		Name, W1, W2, Excepted string
	}{
		{"same length", "abc", "pqr", "apbqcr"},
		{"W2 longer", "ab", "pqrs", "apbqrs"},
		{"W1 longer", "zxcv", "as", "zaxscv"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, mergeAlternately(c.W1, c.W2))
		})
	}
}

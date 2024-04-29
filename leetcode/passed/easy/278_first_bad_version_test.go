package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var isBadVersion func(int) bool

func firstBadVersion(n int) int {
	if n == 1 || isBadVersion(1) {
		return 1
	}

	l, r := 1, n

	for {
		mid := (r-l)/2 + l
		prev := max(0, mid-1)
		isBad := isBadVersion(mid)

		if isBad && !isBadVersion(prev) {
			return mid
		} else if isBad {
			// first bad version is in left part
			r = mid - 1
		} else {
			// first bad version is in right part
			l = mid + 1
		}
	}
}

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		N, Bad int
	}{
		{5, 4},
		{1, 1},
		{2, 1},
		{2, 2},
	}

	for _, c := range cases {
		isBadVersion = func(version int) bool {
			return version >= c.Bad
		}

		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Bad, firstBadVersion(c.N))
		})
	}
}

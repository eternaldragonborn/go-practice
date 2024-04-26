package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	availablePlace := 0
	for index, bed := range flowerbed {
		if bed == 0 {
			prev, next := 0, 0
			if index != 0 {
				prev = flowerbed[index-1]
			}
			if index < len(flowerbed)-1 {
				next = flowerbed[index+1]
			}

			if prev == 0 && next == 0 {
				availablePlace++
				flowerbed[index] = 1
			}
		}
	}

	return availablePlace >= n
}

func TestPlaceFlower(t *testing.T) {
	cases := []struct {
		Flowerbed []int
		N         int
		Excepted  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 0, 0, 1}, 2, true},
		{[]int{0, 0, 0, 0, 1}, 3, false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, canPlaceFlowers(c.Flowerbed, c.N))
		})
	}
}

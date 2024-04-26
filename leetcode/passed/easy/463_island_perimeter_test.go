package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func islandPerimeter(grid [][]int) int {
	ans := 0
	for i, row := range grid {
		for j, col := range row {
			landAround := 0
			if col == 1 {
				// up
				if i > 0 && grid[i-1][j] == 1 {
					landAround++
				}
				// down
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					landAround++
				}
				// left
				if j > 0 && grid[i][j-1] == 1 {
					landAround++
				}
				// right
				if j+1 < len(row) && grid[i][j+1] == 1 {
					landAround++
				}

				ans += (4 - landAround)
			}
		}
	}

	return ans
}

func TestIslandPerimeter(t *testing.T) {
	cases := []struct {
		Grid     [][]int
		Excepted int
	}{
		{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, 16},
		{[][]int{{1}}, 4},
		{[][]int{{1, 0}}, 4},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, islandPerimeter(c.Grid))
		})
	}
}

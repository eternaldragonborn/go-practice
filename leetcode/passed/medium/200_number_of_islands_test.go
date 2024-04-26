package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numIslands(grid [][]byte) int {
	island := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
			return
		}

		grid[i][j] = '0'
		// dfs(i-1, j)
		dfs(i+1, j)
		// dfs(i, j-1)
		dfs(i, j+1)
	}

	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				island++
				dfs(i, j)
			}
		}
	}

	return island
}

func TestNumIsland(t *testing.T) {
	cases := []struct {
		Grid     [][]byte
		Excepted int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, 3},
		{[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, 1},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, numIslands(c.Grid))
		})
	}
}

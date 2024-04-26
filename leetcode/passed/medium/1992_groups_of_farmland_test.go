package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findFarmland(land [][]int) [][]int {
	farmlands := [][]int{}
	rows := len(land)
	cols := len(land[0])

	var dfs func(i, j int, pos *[]int)
	dfs = func(i, j int, pos *[]int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || land[i][j] != 1 {
			return
		}

		land[i][j] = 0
		if i > (*pos)[2] {
			(*pos)[2] = i
		}
		if j > (*pos)[3] {
			(*pos)[3] = j
		}

		dfs(i-1, j, pos)
		dfs(i+1, j, pos)
		dfs(i, j-1, pos)
		dfs(i, j+1, pos)
	}

	for i, row := range land {
		for j, col := range row {
			if col == 1 {
				pos := make([]int, 4)
				pos[0] = i
				pos[1] = j
				dfs(i, j, &pos)
				// fmt.Println(pos)
				farmlands = append(farmlands, pos)
			}
		}
	}
	return farmlands
}

func TestFindFarmLand(t *testing.T) {
	cases := []struct {
		Land, Excepted [][]int
	}{
		{[][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}, [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}}},
		{[][]int{{1, 1}, {1, 1}}, [][]int{{0, 0, 1, 1}}},
		{[][]int{{0}}, [][]int{}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, findFarmland(c.Land))
		})
	}
}

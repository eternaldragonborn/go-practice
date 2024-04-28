package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pixel struct {
	row, col int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// type void struct{}
	m, n := len(image), len(image[0])
	pixels := []pixel{{sr, sc}}
	// visited := make(map[int]void, m*n)
	old_color := image[sr][sc]
	if old_color == color {
		return image
	}

	for len(pixels) != 0 {
		p := pixels[0]
		pixels = pixels[1:]
		// if _, hasVisited := visited[p.row*m+p.col]; hasVisited {
		// 	continue
		// }
		image[p.row][p.col] = color

		// fmt.Println(p.row, p.col)

		up := max(0, p.row-1)
		down := min(m-1, p.row+1)
		left := max(0, p.col-1)
		right := min(n-1, p.col+1)

		if image[up][p.col] == old_color {
			pixels = append(pixels, pixel{up, p.col})
		}
		if image[down][p.col] == old_color {
			pixels = append(pixels, pixel{down, p.col})
		}
		if image[p.row][left] == old_color {
			pixels = append(pixels, pixel{p.row, left})
		}
		if image[p.row][right] == old_color {
			pixels = append(pixels, pixel{p.row, right})
		}
	}

	return image
}

func TestFloodFill(t *testing.T) {
	cases := []struct {
		Image         [][]int
		Sr, Sc, Color int
		Excepted      [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, floodFill(c.Image, c.Sr, c.Sc, c.Color))
		})
	}
}

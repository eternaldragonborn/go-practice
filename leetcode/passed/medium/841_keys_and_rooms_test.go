package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	type void struct{}
	visited := make(map[int]void, len(rooms))

	var dfs func(int, []int)
	dfs = func(room int, keys []int) {
		visited[room] = void{}

		for _, key := range keys {
			if _, exist := visited[key]; !exist {
				dfs(key, rooms[key])
			}
		}
	}
	dfs(0, rooms[0])
	// fmt.Println(visited)

	return len(visited) == len(rooms)
}

func TestVisitAllRooms(t *testing.T) {
	cases := []struct {
		Rooms    [][]int
		Excepted bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, canVisitAllRooms(c.Rooms))
		})
	}
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validPath_dfs(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	type void struct{}
	hasPath := false
	// table of each node can go
	graph := make([][]int, n)
	for _, path := range edges {
		from, to := path[0], path[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	// fmt.Println(graph)

	visited := make(map[int]void)
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = void{}
		for _, to := range graph[from] {
			if to == destination {
				hasPath = true
				return
			}

			if _, hasVisited := visited[to]; !hasVisited {
				dfs(to)
			}
		}
	}

	dfs(source)

	// fmt.Println(visited)
	return hasPath
}

func ValidPath_bfs(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// type void struct{}
	graph := make([][]int, n)
	for _, path := range edges {
		from, to := path[0], path[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	visited := make([]bool, n)
	queue := []int{source}
	for len(queue) != 0 {
		from := queue[0]
		queue = queue[1:]
		if visited[from] {
			continue
		}
		visited[from] = true

		for _, to := range graph[from] {
			if to == destination {
				return true
			}

			if !visited[to] {
				queue = append(queue, to)
			}
		}
	}

	return false
}

func testValidPath(t *testing.T, f func(int, [][]int, int, int) bool) {
	t.Helper()
	cases := []struct {
		N                   int
		Edged               [][]int
		Source, Destination int
		Excepted            bool
	}{
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
		{6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
		{1, [][]int{}, 0, 0, true},
		{5, [][]int{}, 2, 2, true},
		{10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5, true},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, f(c.N, c.Edged, c.Source, c.Destination))
		})
	}
}

func TestValidPath_bfs(t *testing.T) {
	testValidPath(t, ValidPath_bfs)
}

func TestValidPath_dfs(t *testing.T) {
	testValidPath(t, validPath_dfs)
}

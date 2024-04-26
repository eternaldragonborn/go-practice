package leetcode

import (
	"slices"
	"testing"
)

func merge_map(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	// O(N*logN)
	slices.SortFunc(intervals, func(a, b []int) int { return slices.Compare(a, b) })
	positions := make([][]int, 1, len(intervals))
	positions[0] = intervals[0]

	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]
		// O(N)
		if top := positions[len(positions)-1]; start > top[1] {
			// no overlap(start > end of top)
			positions = append(positions, interval)
		} else {
			positions[len(positions)-1][0] = min(start, top[0])
			positions[len(positions)-1][1] = max(end, top[1])
		}
		// O(N^2)
		/* delStart, delLen := -1, 0
		flag := false
		start, end := interval[0], interval[1]

		for index, pos := range positions {
			// overlap
			if start >= pos[0] && start <= pos[1] {
				flag = true
				if delStart == -1 {
					delStart = index
				}
				start = min(start, pos[0])
				end = max(end, pos[1])
				delLen++
			}
		}
		if !flag {
			// no overlap
			positions = append(positions, interval)
		} else {
			if delLen > 1 {
				positions = slices.Delete(positions, delStart+1, delStart+delLen)
			}
			positions[delStart] = []int{start, end}
		} */
	}

	return positions
}

func TestMap(t *testing.T) {
	cases := []struct {
		Intervals [][]int
		Excepted  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 5}}, [][]int{{0, 5}}},
		{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}}},
		{[][]int{{2, 3}, {4, 6}, {5, 7}, {3, 4}}, [][]int{{2, 7}}},
		{[][]int{{2, 3}, {5, 7}, {0, 1}, {4, 6}, {5, 5}, {4, 6}, {5, 6}, {1, 3}, {4, 4}, {0, 0}, {3, 5}, {2, 2}}, [][]int{{0, 7}}},
		{[][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}, [][]int{{1, 3}, {4, 7}}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if ans := merge_map(c.Intervals); !slices.EqualFunc(ans, c.Excepted, func(i1, i2 []int) bool {
				return slices.Equal(i1, i2)
			}) {
				t.Errorf("excepted %v, got %v", c.Excepted, ans)
			}
		})
	}
}

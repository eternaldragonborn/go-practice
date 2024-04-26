package leetcode

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func openLock(deadends []string, target string) int {
	if slices.Contains(deadends, "0000") {
		return -1
	}
	type void struct{}
	lockQueue := []string{"0000"}
	visited := map[string]void{"0000": {}}
	step := -1

	nextLock := func(lock string) []string {
		result := make([]string, 0, 8)

		for index, wheel := range lock {
			var newLock = [2]strings.Builder{}
			for i := 0; i < 4; i++ {
				if i == index {
					if wheel == '0' {
						newLock[0].WriteByte('9')
					} else {
						newLock[0].WriteRune(wheel - 1)
					}

					if wheel == '9' {
						newLock[1].WriteByte('0')
					} else {
						newLock[1].WriteRune(wheel + 1)
					}
				} else {
					newLock[0].WriteByte(lock[i])
					newLock[1].WriteByte(lock[i])
				}
			}

			for _, v := range newLock {
				newLock := v.String()
				if _, isVisited := visited[newLock]; !isVisited && !slices.Contains(deadends, newLock) {
					result = append(result, newLock)
					visited[newLock] = void{}
				}
			}
		}

		return result
	}

	for len(lockQueue) != 0 {
		newQueue := make([]string, 0, len(lockQueue)*8)
		step++
		for _, lock := range lockQueue {
			if lock == target {
				return step
			}

			// push possible next lock into new queue
			newQueue = append(newQueue, nextLock(lock)...)
			// fmt.Println(lock, "->", )

			lockQueue = newQueue
		}
		// fmt.Println(lockQueue, len(lockQueue))
	}
	return -1
}

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "can reach",
			args: args{
				deadends: []string{"0201", "0101", "0102", "1212", "2002"},
				target:   "0202",
			},
			want: 6,
		},
		{
			name: "one move",
			args: args{
				deadends: []string{"8888"},
				target:   "0009",
			},
			want: 1,
		},
		{
			name: "cannot reach",
			args: args{
				deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
				target:   "8888",
			},
			want: -1,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, openLock(tt.args.deadends, tt.args.target))
			// if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
			// 	t.Errorf("openLock() = %v, want %v", got, tt.want)
			// }
		})
	}
}

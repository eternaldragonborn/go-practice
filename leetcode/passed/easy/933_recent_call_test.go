package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RecentCounter struct {
	pingQueue []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.pingQueue) != 0 && this.pingQueue[0] < t-3000 {
		// pop from queue
		this.pingQueue[0] = 0
		this.pingQueue = this.pingQueue[1:]
	}

	this.pingQueue = append(this.pingQueue, t)

	return len(this.pingQueue)
}

func TestRecentCount_933(t *testing.T) {
	cases := []struct {
		T        int
		Excepted int
	}{
		{1, 1},
		{100, 2},
		{3001, 3},
		{3002, 3},
	}

	counter := Constructor()
	for _, c := range cases {
		t.Run(fmt.Sprintf("ping_%d", c.T), func(t *testing.T) {
			assert.Equal(t, c.Excepted, counter.Ping(c.T), counter.pingQueue)
		})
	}
}

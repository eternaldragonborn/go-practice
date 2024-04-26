package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func timeRequiredToBuy(tickets []int, k int) int {
	sum := 0
	n := tickets[k]
	for index, x := range tickets {
		if x < n {
			sum += x
		} else if index <= k {
			sum += n
		} else {
			sum += (n - 1)
		}
	}

	return sum
}

type Case_2073 struct {
	Name        string
	Tickets     []int
	K, Excepted int
}

func TestTimeForTicket(t *testing.T) {
	cases := []Case_2073{
		{
			Name:     "the most",
			Tickets:  []int{5, 1, 1, 1},
			K:        0,
			Excepted: 8,
		},
		{
			Name:     "mid",
			Tickets:  []int{1, 2, 3, 4, 5},
			K:        2,
			Excepted: 10,
		},
		{
			Name:     "mid",
			Tickets:  []int{5, 4, 3, 2, 1},
			K:        2,
			Excepted: 12,
		},
		{
			Name:     "equal",
			Tickets:  []int{2, 3, 2},
			K:        2,
			Excepted: 6,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, timeRequiredToBuy(c.Tickets, c.K))
		})
	}
}

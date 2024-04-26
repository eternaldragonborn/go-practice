package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deckRevealedIncreasing(deck []int) []int {
	slices.Sort(deck)
	if len(deck) == 1 || len(deck) == 2 {
		return deck
	}

	ans, queue := make([]int, len(deck)), make([]int, len(deck))

	// init the queue
	for i := range deck {
		queue[i] = i
	}

	i := 0
	for len(queue) > 0 {
		ans[queue[0]] = deck[i]
		i++
		// GC
		queue[0] = 0
		// remove first element from queue
		if len(queue) == 1 {
			break
		}
		queue = queue[1:]
		// move the second element to last of the queue
		temp := queue[0]
		copy(queue, queue[1:])
		queue[len(queue)-1] = temp
	}

	return ans
}

func TestDeckRevealing(t *testing.T) {
	cases := []struct {
		Name           string
		Deck, Excepted []int
	}{
		{"odd", []int{17, 13, 11, 2, 3, 5, 7}, []int{2, 13, 3, 11, 5, 17, 7}},
		{"even", []int{17, 13, 11, 2, 3, 7}, []int{2, 11, 3, 17, 7, 13}},
		{"two numbers", []int{1, 1000}, []int{1, 1000}},
		{"two numbers", []int{1000, 1}, []int{1, 1000}},
		{"one number", []int{1}, []int{1}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, deckRevealedIncreasing(c.Deck))
		})
	}
}

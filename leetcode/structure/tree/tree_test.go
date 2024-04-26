package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTree(t *testing.T) {
	cases := []struct {
		Name     string
		Data     []any
		Excepted []int
	}{
		{"empty", []any{}, []int{}},
		{"no nil nodes", []any{3, 9, 20, 7, 10, 15, 7}, []int{3, 9, 20, 7, 10, 15, 7}},
		{"have nil nodes", []any{3, 9, 20, nil, nil, 15, 7}, []int{3, 9, 20, 15, 7}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, root.LevelOrderTraversal())
		})
	}
}

func TestMaximumDepth(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted int
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, 3},
		{[]any{1, nil, 2}, 2},
		{[]any{1, 2, 2, 3, 3, nil, nil, 4, 4}, 4},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, treeHeight(root))
		})
	}
}

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		Data     []any
		Excepted bool
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, true},
		{[]any{1, 2, 2, 3, 3, nil, nil, 4, 4}, false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			assert.Equal(t, c.Excepted, isBalanced(root))
		})
	}
}

func TestPreOrder(t *testing.T) {
	cases := []struct {
		Data []any
	}{
		{[]any{4, 1, 1, 2, nil, nil, 6, 3, 1, 5}},
		{[]any{4, 2, nil, 1, 1, 3, nil, nil, 1}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			root := CreateTree(c.Data)
			ans := root.PreOrderTraversal()
			assert.Equal(t, c.Data, ans)
		})
	}
}

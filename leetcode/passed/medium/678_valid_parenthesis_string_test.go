package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkValidString(s string) bool {
	leftParenthesis, stars := []int8{}, []int8{}

	for index, c := range s {
		index := int8(index)
		if c == '(' {
			leftParenthesis = append(leftParenthesis, index)
		} else if c == '*' {
			stars = append(stars, index)
		} else { // ')'
			if len(leftParenthesis) == 0 && len(stars) == 0 { // invalid
				return false
			}

			if len(leftParenthesis) != 0 {
				// GC
				leftParenthesis[len(leftParenthesis)-1] = 0
				// pop '('
				leftParenthesis = leftParenthesis[:len(leftParenthesis)-1]
			} else {
				stars[len(stars)-1] = 0
				// pop '*'
				stars = stars[:len(stars)-1]
			}
		}
	}

	// fmt.Println(leftParenthesis, stars)

	if len(leftParenthesis) == 0 {
		return true
	}

	for i := len(leftParenthesis) - 1; i >= 0; i-- {
		// no '*' or stars are before '('(redundant)
		if len(stars) == 0 || leftParenthesis[i] > stars[len(stars)-1] {
			return false
		}

		stars[len(stars)-1] = 0
		// pop '*'
		stars = stars[:len(stars)-1]
	}

	return true
}

func TestParenthesis_678(t *testing.T) {
	cases := []struct {
		S        string
		Excepted bool
	}{
		{"()", true},
		{"(*)", true},
		{"(*))", true},
		{"**********", true},
		{"((*", false},
		{"*)))", false},
		{"(()*()))", true},
		{"**(", false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.Excepted, checkValidString(c.S), c.S)
		})
	}
}

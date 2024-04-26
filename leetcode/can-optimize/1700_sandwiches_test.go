package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 43.04% / 40.51%
func countStudents(students []int, sandwiches []int) int {
	studentCounting := [2]int{0, 0}
	for len(students) != 0 {
		stu := students[0]
		if studentCounting[0]+studentCounting[1] != len(students) {
			studentCounting[stu]++
		}

		if stu == sandwiches[0] {
			// pop student
			students[0] = 0
			students = students[1:]
			studentCounting[stu]--

			// pop sandwich
			sandwiches[0] = 0
			sandwiches = sandwiches[1:]
		} else if studentCounting[0]+studentCounting[1] == len(students) &&
			(studentCounting[0] == 0 || studentCounting[1] == 0) {
			return len(students)
		} else {
			// push the student to end of queue
			copy(students, students[1:])
			students[len(students)-1] = stu
			// students = append(students[1:], stu)
		}
	}

	return 0
}

func TestStudentsLeft(t *testing.T) {
	cases := []struct {
		Name                 string
		Students, Sandwiches []int
		Excepted             int
	}{
		{"not students left", []int{1, 1, 0, 0}, []int{0, 1, 0, 1}, 0},
		{"left", []int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}, 3},
		{"", []int{0, 0, 0, 1, 1, 1, 1, 0, 0, 0}, []int{1, 0, 1, 0, 0, 1, 1, 0, 0, 0}, 0},
		{"", []int{0, 0, 0, 1, 1, 1, 1, 0, 0, 0}, []int{1, 0, 1, 0, 0, 0, 1, 0, 0, 0}, 1},
		//
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.Excepted, countStudents(c.Students, c.Sandwiches))
		})
	}
}

package day4

import (
	"strconv"
	"strings"
)

type Assignment struct {
	Left  int
	Right int
}

func NewAssignment(assignmentStr string) *Assignment {
	assignment := strings.Split(assignmentStr, "-")

	left, err := strconv.Atoi(assignment[0])
	if err != nil {
		panic(err)
	}

	right, err := strconv.Atoi(assignment[1])
	if err != nil {
		panic(err)
	}

	return &Assignment{
		Left:  left,
		Right: right,
	}
}

func (A *Assignment) Overlaps(candidate *Assignment) bool {
	return (A.Left >= candidate.Left && A.Left <= candidate.Right) ||
		(A.Right >= candidate.Left && A.Right <= candidate.Right) ||
		(candidate.Left >= A.Left && candidate.Left <= A.Right) ||
		(candidate.Right >= A.Left && candidate.Right <= A.Right)
}

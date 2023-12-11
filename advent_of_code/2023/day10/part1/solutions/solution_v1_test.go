package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day10/part1/solutions"
)

func TestSolver(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			solver := solutions.NewSolver()
			solver.LoadArrayOfLines(input)
			solver.ComputeResult()

			assert.Equal(t, tc.expectedResult, solver.Result())
		})
	}
}

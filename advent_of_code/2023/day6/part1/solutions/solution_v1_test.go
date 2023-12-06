package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day6/part1/solutions"
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

func TestCountHowManyWaysOfWinning(t *testing.T) {
	testCases := []struct {
		name           string
		raceDuration   int
		bestDistance   int
		expectedResult int
	}{
		{
			name:           "test 1",
			raceDuration:   7,
			bestDistance:   9,
			expectedResult: 4,
		},
		{
			name:           "test 2",
			raceDuration:   15,
			bestDistance:   40,
			expectedResult: 8,
		},
		{
			name:           "test 3",
			raceDuration:   30,
			bestDistance:   200,
			expectedResult: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.CountHowManyWaysOfWinning(tc.raceDuration, tc.bestDistance)
			assert.Equal(t, tc.expectedResult, result)

		})
	}
}

package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day3/part1/solutions"
)

func TestComputePartNumber(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			partNumberFinder := solutions.NewPartNumberFinder()
			partNumberFinder.LoadArrayOfLines(input)

			err := partNumberFinder.ComputePartNumber()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, partNumberFinder.PartNumber())
		})
	}
}

func TestFindNumbersInLine(t *testing.T) {
	testCases := []struct {
		name           string
		line           string
		expectedResult []solutions.NumberIndexes
	}{
		{
			name: "test 1",
			line: "467..114..",
			expectedResult: []solutions.NumberIndexes{
				{
					Start: 0,
					Stop:  2,
				},
				{
					Start: 5,
					Stop:  7,
				},
			},
		},
		{
			name: "test 2",
			line: "467....155",
			expectedResult: []solutions.NumberIndexes{
				{
					Start: 0,
					Stop:  2,
				},
				{
					Start: 7,
					Stop:  9,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.FindNumbersInLine(tc.line)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestGetNumberSurroundingLimits(t *testing.T) {
	testCases := []struct {
		name                string
		nrIdx               solutions.NumberIndexes
		cols                int
		expectedResultStart int
		expectedResultStop  int
	}{
		{
			name: "test 1",
			nrIdx: solutions.NumberIndexes{
				Start: 0,
				Stop:  2,
			},
			cols:                4,
			expectedResultStart: 0,
			expectedResultStop:  3,
		},
		{
			name: "test 2",
			nrIdx: solutions.NumberIndexes{
				Start: 2,
				Stop:  5,
			},
			cols:                6,
			expectedResultStart: 1,
			expectedResultStop:  5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultStart, resultStop := solutions.GetNumberSurroundingLimits(tc.nrIdx, tc.cols)
			assert.Equal(t, tc.expectedResultStart, resultStart)
			assert.Equal(t, tc.expectedResultStop, resultStop)
		})
	}
}

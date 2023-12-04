package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day4/part1/solutions"
)

func TestScratchCardValidator(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			scratchCardValidator := solutions.NewScratchCardValidator()

			err := scratchCardValidator.LoadArrayOfLines(input)
			require.NoError(t, err)

			assert.Equal(t, tc.expectedResult, scratchCardValidator.PointsAccumulator())
		})
	}
}

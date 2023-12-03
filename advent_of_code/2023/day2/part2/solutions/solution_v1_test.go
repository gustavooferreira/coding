package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day2/part2/solutions"
)

func TestComputeMinimumGameSetArrayOfLines(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			powerCalculator := solutions.NewPowerCalculator()

			err := powerCalculator.ComputeMinimumGameSetArrayOfLines(input)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, powerCalculator.GameSetPowerAccumulator())
		})
	}
}

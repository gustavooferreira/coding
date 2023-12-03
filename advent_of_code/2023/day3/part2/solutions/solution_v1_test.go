package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day3/part2/solutions"
)

func TestComputeGearRatioSum(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			gearRatioFinder := solutions.NewGearRatioFinder()
			gearRatioFinder.LoadArrayOfLines(input)

			err := gearRatioFinder.ComputeGearRatioSum()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, gearRatioFinder.GearRatioSum())
		})
	}
}

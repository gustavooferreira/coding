package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day1/part2/solutions"
)

func TestMyFunc(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			calibrator := solutions.NewCalibrator()

			calibrator.CalculateCalibrationArrayOfLines(input)
			assert.Equal(t, tc.expectedResult, calibrator.Accumulator())
		})
	}
}

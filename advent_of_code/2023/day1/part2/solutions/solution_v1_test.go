package solutions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent_of_code/2023/day1/part2/solutions"
)

func TestMyFunc(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Split(tc.text, "\n")

			calibrator := solutions.NewCalibrator()

			calibrator.CalculateCalibrationArrayOfLines(input)
			assert.Equal(t, tc.expectedResult, calibrator.GetAccumulator())
		})
	}
}

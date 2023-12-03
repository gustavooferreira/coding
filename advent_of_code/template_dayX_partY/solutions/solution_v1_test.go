package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/template_dayX_partY/solutions"
)

func TestCalculateCalibrationArrayOfLines(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			calibrator := solutions.NewSolver()

			calibrator.LoadArrayOfLines(input)
			assert.Equal(t, tc.expectedResult, calibrator.Accumulator())
		})
	}
}

package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent_of_code/2023/day1/part2/solutions"
)

func TestMyFunc(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.CalculateCalibration(tc.text)
			assert.Equal(t, tc.expectedResult, result)
		})
	}

}

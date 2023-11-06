package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"problem_solving/1_two_sum/solutions"
)

func TestTwoSum_V1(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.TwoSum_V1(tc.nums, tc.target)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

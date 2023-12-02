package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/leetcode/1/solutions"
)

func TestTwoSum_V3(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.TwoSum_V3(tc.nums, tc.target)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

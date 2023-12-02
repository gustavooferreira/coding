package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/leetcode/3/solutions"
)

func TestLengthOfLongestSubstrings_V2(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.LengthOfLongestSubstring_V1(tc.s)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

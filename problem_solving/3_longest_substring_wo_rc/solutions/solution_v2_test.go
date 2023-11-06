package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"problem_solving/3_longest_substring_wo_rc/solutions"
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

package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/leetcode/3/solutions"
)

func TestLengthOfLongestSubstrings_V1(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.LengthOfLongestSubstring_V1(tc.s)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestCountSequenceOfNonRepeatingCharacters(t *testing.T) {
	result := solutions.CountSequenceOfNonRepeatingCharacters("abceaklui")
	assert.Equal(t, 4, result)
}

package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/leetcode/2/solutions"
)

func TestAddTwoNumbers_V1(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := solutions.CreateLinkedListFromSlice(tc.l1Slice)
			l2 := solutions.CreateLinkedListFromSlice(tc.l2Slice)

			result := solutions.AddTwoNumbers_V1(l1, l2)

			resultSlice := solutions.GetSliceFromLinkedList(result)
			assert.Equal(t, tc.expectedResult, resultSlice)
		})
	}
}

package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"problem_solving/2_add_two_numbers/solutions"
)

func TestCreateLinkedListFromSlice(t *testing.T) {
	testCases := []struct {
		name           string
		val            []int
		expectedResult []int
	}{
		{
			name:           "test nil slices",
			val:            nil,
			expectedResult: nil,
		},
		{
			name:           "test empty slices",
			val:            []int{},
			expectedResult: nil,
		},
		{
			name:           "test one digit only",
			val:            []int{5},
			expectedResult: []int{5},
		},
		{
			name:           "test 123",
			val:            []int{1, 2, 3},
			expectedResult: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.CreateLinkedListFromSlice(tc.val)
			resultSlice := solutions.GetSliceFromLinkedList(result)
			assert.Equal(t, tc.expectedResult, resultSlice)

		})
	}
}

func TestIntToSlice(t *testing.T) {
	testCases := []struct {
		name           string
		num            int
		expectedResult []int
	}{
		{
			name:           "test 123",
			num:            123,
			expectedResult: []int{1, 2, 3},
		},
		{
			name:           "test 654",
			num:            654,
			expectedResult: []int{6, 5, 4},
		},
		{
			name:           "test 100",
			num:            100,
			expectedResult: []int{1, 0, 0},
		},
		{
			name:           "test 12121",
			num:            12121,
			expectedResult: []int{1, 2, 1, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.IntToSlice(tc.num)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestCompareLinkedLists(t *testing.T) {
	testCases := []struct {
		name           string
		val1           []int
		val2           []int
		expectedResult bool
	}{
		{
			name:           "test nil slices",
			val1:           nil,
			val2:           nil,
			expectedResult: true,
		},
		{
			name:           "test empty slices",
			val1:           []int{},
			val2:           []int{},
			expectedResult: true,
		},
		{
			name:           "test equal slices with one digit only",
			val1:           []int{5},
			val2:           []int{5},
			expectedResult: true,
		},
		{
			name:           "test equal slices",
			val1:           []int{1, 2, 3},
			val2:           []int{1, 2, 3},
			expectedResult: true,
		},
		{
			name:           "test non equal slices",
			val1:           []int{1, 2, 3},
			val2:           []int{3, 5, 4},
			expectedResult: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := solutions.CreateLinkedListFromSlice(tc.val1)
			l2 := solutions.CreateLinkedListFromSlice(tc.val2)

			result := solutions.CompareLinkedLists(l1, l2)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

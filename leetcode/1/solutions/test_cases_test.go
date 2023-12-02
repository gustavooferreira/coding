package solutions_test

type testCase struct {
	name           string
	nums           []int
	target         int
	expectedResult []int
}

func getTestCases() []testCase {
	return []testCase{
		{
			name:           "test with zero elements in nums slice",
			nums:           []int{},
			target:         10,
			expectedResult: nil,
		},
		{
			name:           "test with one element in nums slice",
			nums:           []int{5},
			target:         10,
			expectedResult: nil,
		},
		{
			name:           "test 1",
			nums:           []int{2, 7, 11, 15},
			target:         9,
			expectedResult: []int{0, 1},
		},
		{
			name:           "test 2",
			nums:           []int{3, 2, 4},
			target:         6,
			expectedResult: []int{1, 2},
		},
		{
			name:           "test 3",
			nums:           []int{3, 3},
			target:         6,
			expectedResult: []int{0, 1},
		},
		{
			name:           "test 4 - contains the same values multiple times",
			nums:           []int{7, 6, 3, 7, 9, 1},
			target:         14,
			expectedResult: []int{0, 3},
		},
	}
}

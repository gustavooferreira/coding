package solutions_test

type testCase struct {
	name           string
	l1Slice        []int
	l2Slice        []int
	expectedResult []int
}

func getTestCases() []testCase {
	return []testCase{
		{
			name:           "test nil slices",
			l1Slice:        nil,
			l2Slice:        nil,
			expectedResult: nil,
		},
		{
			name:           "test correct sum 342 + 465 = 807",
			l1Slice:        []int{3, 4, 2},
			l2Slice:        []int{4, 6, 5},
			expectedResult: []int{8, 0, 7},
		},
		{
			name:           "test correct sum 0 + 0 = 0",
			l1Slice:        []int{0},
			l2Slice:        []int{0},
			expectedResult: []int{0},
		},
		{
			name:           "test correct sum 99999999 + 9999 = 10009998",
			l1Slice:        []int{9, 9, 9, 9, 9, 9, 9},
			l2Slice:        []int{9, 9, 9, 9},
			expectedResult: []int{1, 0, 0, 0, 9, 9, 9, 8},
		},
		{
			name:           "test correct sum 120 + 320 = 440",
			l1Slice:        []int{1, 2, 0},
			l2Slice:        []int{3, 2, 0},
			expectedResult: []int{4, 4, 0},
		},
		{
			name:           "test correct sum [10,20,30] + [30,40,50] = 4680",
			l1Slice:        []int{10, 20, 30},
			l2Slice:        []int{30, 40, 50},
			expectedResult: []int{4, 6, 8, 0},
		},
	}
}

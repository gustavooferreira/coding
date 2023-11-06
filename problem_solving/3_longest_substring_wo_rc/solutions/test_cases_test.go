package solutions_test

type testCase struct {
	name           string
	s              string
	expectedResult int
}

func getTestCases() []testCase {
	return []testCase{
		{
			name:           "test empty string",
			s:              "",
			expectedResult: 0,
		},
		{
			name:           "test string 'abcabcbb'",
			s:              "abcabcbb",
			expectedResult: 3,
		},
		{
			name:           "test string 'bbbbb'",
			s:              "bbbbb",
			expectedResult: 1,
		},
		{
			name:           "test string 'pwwkew'",
			s:              "pwwkew",
			expectedResult: 3,
		},
	}
}

package solutions_test

import "github.com/gustavooferreira/coding/advent_of_code/internal/testutils"

type testCase struct {
	name           string
	textInfo       testutils.TextInfo
	expectedResult int
}

func getTestCases() []testCase {
	return []testCase{
		{
			name: "problem statement example",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1.txt",
			},
			expectedResult: 82000210,
		},
		{
			name: "problem statement challenge",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input2.txt",
			},
			expectedResult: 382979724122,
		},
	}
}

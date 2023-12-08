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
			name: "problem statement example A",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1a.txt",
			},
			expectedResult: 2,
		},
		{
			name: "problem statement example B",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1b.txt",
			},
			expectedResult: 6,
		},
		{
			name: "problem statement challenge",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input2.txt",
			},
			expectedResult: 16531,
		},
	}
}

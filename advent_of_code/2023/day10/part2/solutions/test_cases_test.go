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
			name: "problem statement example 1a",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1a.txt",
			},
			expectedResult: 4,
		},
		{
			name: "problem statement example 1b",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1b.txt",
			},
			expectedResult: 8,
		},
		{
			name: "problem statement example 1c",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1c.txt",
			},
			expectedResult: 10,
		},
		{
			name: "problem statement challenge",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input2.txt",
			},
			expectedResult: 0,
		},
	}
}

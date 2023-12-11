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
			name: "problem statement example a1",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1a1.txt",
			},
			expectedResult: 4,
		},
		{
			name: "problem statement example a2",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1a2.txt",
			},
			expectedResult: 4,
		},
		{
			name: "problem statement example b1",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1b1.txt",
			},
			expectedResult: 8,
		},
		{
			name: "problem statement example b2",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1b2.txt",
			},
			expectedResult: 8,
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

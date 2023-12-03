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
			name: "no input text",
			textInfo: testutils.TextInfo{
				Content: "",
			},
			expectedResult: 0,
		},
		{
			name: "problem statement example",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input1.txt",
			},
			expectedResult: 142,
		},
		{
			name: "problem statement challenge",
			textInfo: testutils.TextInfo{
				FilePath: "testdata/input2.txt",
			},
			expectedResult: 56397,
		},
		{
			name: "example with empty lines",
			textInfo: testutils.TextInfo{
				Content: `
pqr3stu8vwx

treb7uc0het

`,
			},
			expectedResult: 108,
		},
	}
}

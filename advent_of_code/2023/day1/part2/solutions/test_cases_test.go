package solutions_test

type testCase struct {
	name           string
	text           string
	expectedResult int
}

func getTestCases() []testCase {
	return []testCase{
		{
			name:           "no input text",
			text:           "",
			expectedResult: 0,
		},
		{
			name: "problem statement example",
			text: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			expectedResult: 281,
		},
		{
			name: "example with empty lines",
			text: `
pqr3stu8vwx


treb7uc0het

`,
			expectedResult: 108,
		},
	}
}

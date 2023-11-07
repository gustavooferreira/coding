package solutions_test

type testCase struct {
	name           string
	str            string
	expectedResult bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			name:           "matching parenthesis should succeed",
			str:            "(this is a phrase {that contains another phrase [123]})",
			expectedResult: true,
		},
		{
			name:           "non matching parenthesis should fail",
			str:            "this [ is not match by the right parenthesis close)",
			expectedResult: false,
		},
		{
			name:           "multiple types of matching parenthesis inside and in sequence should succeed",
			str:            "([{[]}])[]{{()}}",
			expectedResult: true,
		},
		{
			name:           "unclosed parenthesis should fail",
			str:            "(kljkljk[kljkjklj{kjljlkj}] dsfsdfsdf",
			expectedResult: false,
		},
	}
}

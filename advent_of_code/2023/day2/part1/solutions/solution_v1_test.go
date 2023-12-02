package solutions_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day2/part1/solutions"
)

func TestValidateGameArrayOfLines(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Split(tc.text, "\n")

			validator := solutions.NewValidator()

			validator.ValidateGameArrayOfLines(input)
			assert.Equal(t, tc.expectedResult, validator.GetGameIDAccumulator())
		})
	}
}

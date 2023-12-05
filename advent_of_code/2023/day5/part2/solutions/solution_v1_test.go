package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day5/part2/solutions"
)

func TestSolver(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.textInfo.GetTextContent(t)

			solver := solutions.NewSolver()
			solver.LoadArrayOfLines(input)
			solver.ComputeResult()

			assert.Equal(t, tc.expectedResult, solver.Result())
		})
	}
}

func TestGenericMapper(t *testing.T) {
	t.Run("should add ranges in the correct order", func(t *testing.T) {
		genericMapper := solutions.GenericMapper{
			Source:      "source",
			Destination: "destination",
		}

		genericMapper.AddRangeMapping(150, 50, 10)
		genericMapper.AddRangeMapping(180, 80, 10)
		genericMapper.AddRangeMapping(120, 20, 5)

		result := genericMapper.Ranges()

		t.Logf("%+v", result)

		expectedResult := []solutions.RangeMapperInfo{
			{
				DestinationStart: 120,
				SourceStart:      20,
				Length:           5,
			},
			{
				DestinationStart: 150,
				SourceStart:      50,
				Length:           10,
			},
			{
				DestinationStart: 180,
				SourceStart:      80,
				Length:           10,
			},
		}

		assert.Equal(t, expectedResult, result)
	})
}

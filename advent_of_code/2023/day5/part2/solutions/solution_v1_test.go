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

func TestGenericMapperAddition(t *testing.T) {
	t.Run("should add ranges in the correct order", func(t *testing.T) {
		genericMapper := solutions.GenericMapper{
			Source:      "source",
			Destination: "destination",
		}

		genericMapper.AddRangeMapping(150, 50, 10)
		genericMapper.AddRangeMapping(180, 80, 10)
		genericMapper.AddRangeMapping(120, 20, 5)

		result := genericMapper.Ranges()

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

func TestGenericMapperLookup(t *testing.T) {
	t.Run("should perform correct lookups on range that span the entire mapping", func(t *testing.T) {
		genericMapper := solutions.GenericMapper{
			Source:      "source",
			Destination: "destination",
		}

		// Line: 10 -- 20 -- 25 -- 50 -- 60 -- 80 -- 90

		genericMapper.AddRangeMapping(1050, 50, 10)
		genericMapper.AddRangeMapping(1080, 80, 10)
		genericMapper.AddRangeMapping(1020, 20, 5)

		result := genericMapper.Lookup(solutions.ItemRange{
			Start:  10,
			Length: 100,
		})

		t.Logf("%+v", result)

		expectedResult := []solutions.ItemRange{
			{
				Start:  10,
				Length: 10,
			},
			{
				Start:  1020,
				Length: 5,
			},
			{
				Start:  25,
				Length: 25,
			},
			{
				Start:  1050,
				Length: 10,
			},
			{
				Start:  60,
				Length: 20,
			},
			{
				Start:  1080,
				Length: 10,
			},
			{
				Start:  90,
				Length: 20,
			},
		}

		assert.ElementsMatch(t, expectedResult, result)
	})

	t.Run("should perform correct lookups on a range that spans only a portion of the mappings", func(t *testing.T) {
		genericMapper := solutions.GenericMapper{
			Source:      "source",
			Destination: "destination",
		}

		// Line: 25 -- 40 -- 50 -- 54

		genericMapper.AddRangeMapping(1020, 20, 20)
		genericMapper.AddRangeMapping(1050, 50, 10)

		result := genericMapper.Lookup(solutions.ItemRange{
			Start:  25,
			Length: 30,
		})

		t.Logf("%+v", result)

		expectedResult := []solutions.ItemRange{
			{
				Start:  1025,
				Length: 15,
			},
			{
				Start:  40,
				Length: 10,
			},
			{
				Start:  1050,
				Length: 5,
			},
		}

		assert.ElementsMatch(t, expectedResult, result)
	})
}

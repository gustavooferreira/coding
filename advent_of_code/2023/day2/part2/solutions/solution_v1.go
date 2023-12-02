package solutions

import (
	"strconv"
	"strings"
)

type PowerCalculator struct {
	gameSetPowerAccumulator int
}

func NewPowerCalculator() *PowerCalculator {
	return &PowerCalculator{}
}

// ComputeMinimumGameSetArrayOfLines takes an array of lines and adds the cube set power to the accumulator.
func (c *PowerCalculator) ComputeMinimumGameSetArrayOfLines(lines []string) {
	for _, line := range lines {
		c.ComputeMinimumGameSetForLine(line)
	}
}

// ComputeMinimumGameSetForLine finds the minimum set of cubes that must have been present in the bag for the given
// game, and adds its power to the accumulator.
func (c *PowerCalculator) ComputeMinimumGameSetForLine(line string) {
	if len(line) == 0 {
		return
	}

	// split and trim on the `:` and get the game ID
	result := strings.Split(line, ":")

	// split and trim on the `;` to get the various sets.
	resultSetsString := strings.Split(result[1], ";")

	maxRedCubeCount := 0
	maxGreenCubeCount := 0
	maxBlueCubeCount := 0

	// split and trim on the comma to get each cube type result
	// Go over the game set
	for _, resultSetString := range resultSetsString {
		cubesString := strings.Split(resultSetString, ",")

		for _, cubeString := range cubesString {
			cubeString = strings.TrimSpace(cubeString)
			cubeInfoArr := strings.Split(cubeString, " ")
			cubeCount, _ := strconv.Atoi(cubeInfoArr[0])
			cubeColour := cubeInfoArr[1]

			if cubeColour == "red" {
				if cubeCount > maxRedCubeCount {
					maxRedCubeCount = cubeCount
				}
			} else if cubeColour == "green" {
				if cubeCount > maxGreenCubeCount {
					maxGreenCubeCount = cubeCount
				}
			} else if cubeColour == "blue" {
				if cubeCount > maxBlueCubeCount {
					maxBlueCubeCount = cubeCount
				}
			}
		}
	}

	// Calculate set of cubes power
	setPower := maxRedCubeCount * maxGreenCubeCount * maxBlueCubeCount

	c.gameSetPowerAccumulator += setPower
}

// GetGameSetPowerAccumulator returns the current result stored in the game set power accumulator.
func (c *PowerCalculator) GetGameSetPowerAccumulator() int {
	return c.gameSetPowerAccumulator
}

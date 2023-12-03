package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type PowerCalculator struct {
	debug                   bool
	gameSetPowerAccumulator int
}

func NewPowerCalculator() *PowerCalculator {
	return &PowerCalculator{}
}

func (pc *PowerCalculator) SetDebug(enable bool) {
	pc.debug = enable
}

// ComputeMinimumGameSetArrayOfLines takes an array of lines and adds the cube set power to the accumulator.
func (pc *PowerCalculator) ComputeMinimumGameSetArrayOfLines(lines []string) error {
	for i, line := range lines {
		err := pc.ComputeMinimumGameSetForLine(i+1, line)
		if err != nil {
			return fmt.Errorf("error on line %d: %w", i+1, err)
		}
	}

	return nil
}

// ComputeMinimumGameSetForLine finds the minimum set of cubes that must have been present in the bag for the given
// game, and adds its power to the accumulator.
func (pc *PowerCalculator) ComputeMinimumGameSetForLine(lineNumber int, line string) error {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return nil
	}

	// split game number from the cube sets
	result := strings.Split(line, ":")
	if len(result) != 2 {
		return fmt.Errorf("game line expected to contain a single ':' character")
	}

	gameIDString := strings.TrimSpace(result[0])
	gameIDString = strings.TrimPrefix(gameIDString, "Game ")
	gameIDString = strings.TrimSpace(gameIDString)
	gameID, err := strconv.Atoi(gameIDString)
	if err != nil {
		return fmt.Errorf("game ID %q is not valid: %w", gameIDString, err)
	}

	// split game into the various subsets of cubes
	resultSetsString := strings.Split(result[1], ";")

	maxCountCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	// split and trim on the comma to get each cube type result
	// Go over the game set
	for _, resultSetString := range resultSetsString {
		cubesString := strings.Split(resultSetString, ",")

		for _, cubeString := range cubesString {
			cubeString = strings.TrimSpace(cubeString)
			cubeInfoArr := strings.Split(cubeString, " ")
			if len(cubeInfoArr) != 2 {
				return fmt.Errorf("cube info inside cube set expected to contain number of cubes for a given colour separated by white space")
			}

			cubeCount, err := strconv.Atoi(cubeInfoArr[0])
			if err != nil {
				return fmt.Errorf("cube count %q is not valid: %w", cubeCount, err)
			}

			cubeColour := cubeInfoArr[1]

			if cubeCount > maxCountCubes[cubeColour] {
				maxCountCubes[cubeColour] = cubeCount
			}
		}
	}

	if pc.debug {
		fmt.Printf("Line: %3d -- Game ID: %3d -- MaxCounts: {Red: %3d, Green: %3d, Blue: %3d}\n", lineNumber, gameID,
			maxCountCubes["red"], maxCountCubes["green"], maxCountCubes["blue"])
	}

	// Calculate set of cubes power
	setPower := maxCountCubes["red"] * maxCountCubes["green"] * maxCountCubes["blue"]

	pc.gameSetPowerAccumulator += setPower
	return nil
}

// GameSetPowerAccumulator returns the current result stored in the game set power accumulator.
func (pc *PowerCalculator) GameSetPowerAccumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return pc.gameSetPowerAccumulator
}

package solutions

import (
	"strconv"
	"strings"
)

// change for map
const (
	redCubesTotal   int = 12
	greenCubesTotal int = 13
	blueCubesTotal  int = 14
)

type Validator struct {
	gameIDAccumulator int
}

func NewValidator() *Validator {
	return &Validator{}
}

// ValidateGameArrayOfLines takes an array of lines and adds their game ID number to the accumulator if the game is
// valid.
func (c *Validator) ValidateGameArrayOfLines(lines []string) {
	for _, line := range lines {
		c.ValidateGameForLine(line)
	}
}

// ValidateGameForLine validates whether the game represented by the line is valid and if so, add the game ID to the
// accumulator.
func (c *Validator) ValidateGameForLine(line string) {
	if len(line) == 0 {
		return
	}

	// split and trim on the `:` and get the game ID
	result := strings.Split(line, ":")
	gameIDString := strings.TrimSpace(result[0])
	gameIDString = strings.TrimPrefix(gameIDString, "Game ")
	gameIDString = strings.TrimSpace(gameIDString)
	gameID, _ := strconv.Atoi(gameIDString)

	// split and trim on the `;` to get the various sets.
	resultSetsString := strings.Split(result[1], ";")

	// split and trim on the comma to get each cube type result
	for _, resultSetString := range resultSetsString {
		cubesString := strings.Split(resultSetString, ",")

		for _, cubeString := range cubesString {
			cubeString = strings.TrimSpace(cubeString)
			cubeInfoArr := strings.Split(cubeString, " ")
			cubeCount, _ := strconv.Atoi(cubeInfoArr[0])
			cubeColour := cubeInfoArr[1]

			if cubeColour == "red" {
				if cubeCount > redCubesTotal {
					return // invalid game
				}
			} else if cubeColour == "green" {
				if cubeCount > greenCubesTotal {
					return // invalid game
				}
			} else if cubeColour == "blue" {
				if cubeCount > blueCubesTotal {
					return // invalid game
				}
			}
		}
	}

	// if we got here, game is valid
	c.gameIDAccumulator += gameID
}

// GetGameIDAccumulator returns the current result stored in the Game ID accumulator.
func (c *Validator) GetGameIDAccumulator() int {
	return c.gameIDAccumulator
}

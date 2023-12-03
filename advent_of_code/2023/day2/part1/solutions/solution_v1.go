package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// totals represents the total number of cubes allowed in a game.
// These should be treated as constants.
var totals = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Validator struct {
	debug             bool
	gameIDAccumulator int
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) SetDebug(enable bool) {
	v.debug = enable
}

// ValidateGameArrayOfLines takes an array of lines and adds their game ID number to the accumulator if the game is
// valid.
func (v *Validator) ValidateGameArrayOfLines(lines []string) error {
	for i, line := range lines {
		err := v.ValidateGameForLine(i+1, line)
		if err != nil {
			return fmt.Errorf("error on line %d: %w", i+1, err)
		}
	}

	return nil
}

// ValidateGameForLine validates whether the game represented by the line is valid and if so, add the game ID to the
// accumulator.
func (v *Validator) ValidateGameForLine(lineNumber int, line string) error {
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

	gameInfo := GameInfo{
		ID:    gameID,
		Valid: true, // start with the assumption the game is valid
	}

	// split game into the various subsets of cubes
	resultSetsString := strings.Split(result[1], ";")

	for _, resultSetString := range resultSetsString {
		cubesString := strings.Split(resultSetString, ",")

		cubeSet := CubeSet{}

		for _, cubeString := range cubesString {
			cubeString = strings.TrimSpace(cubeString)
			cubeInfoArr := strings.Split(cubeString, " ")
			if len(cubeInfoArr) != 2 {
				return fmt.Errorf("cube info inside cube set expected to contain number of cubes for a given colour separated by white space")
			}

			cubeCount, err := strconv.Atoi(cubeInfoArr[0])
			if err != nil {
				return fmt.Errorf("cube count %q for game ID %d is not valid: %w", cubeCount, gameID, err)
			}
			cubeColour := cubeInfoArr[1]

			cubeSet = SetCubeSetField(cubeSet, cubeColour, cubeCount)

			if cubeCount > totals[cubeColour] {
				gameInfo.Valid = false
			}
		}

		gameInfo.CubeSets = append(gameInfo.CubeSets, cubeSet)
	}

	if v.debug {
		fmt.Printf("Line: %3d -- Game ID: %3d -- Valid: %-5t -- Sets: %+v\n", lineNumber, gameInfo.ID, gameInfo.Valid, gameInfo.CubeSets)
	}

	if !gameInfo.Valid {
		return nil
	}

	// if we got here, game is valid
	v.gameIDAccumulator += gameID
	return nil
}

// GameIDAccumulator returns the current result stored in the Game ID accumulator.
func (v *Validator) GameIDAccumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return v.gameIDAccumulator
}

type GameInfo struct {
	ID       int
	CubeSets []CubeSet
	// Returns whether this game is valid
	Valid bool
}

type CubeSet struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

// ----- Helper functions -----

func SetCubeSetField(cubeSet CubeSet, colour string, count int) CubeSet {
	switch colour {
	case "red":
		cubeSet.RedCount = count
	case "green":
		cubeSet.GreenCount = count
	case "blue":
		cubeSet.BlueCount = count
	}

	return cubeSet
}

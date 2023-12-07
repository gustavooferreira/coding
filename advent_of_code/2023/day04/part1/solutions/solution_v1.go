package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type ScratchCardValidator struct {
	debug             bool
	pointsAccumulator int
}

func NewScratchCardValidator() *ScratchCardValidator {
	return &ScratchCardValidator{}
}

func (scv *ScratchCardValidator) SetDebug(enable bool) {
	scv.debug = enable
}

// LoadArrayOfLines takes an array of lines and adds their value number to the accumulator.
func (scv *ScratchCardValidator) LoadArrayOfLines(lines []string) error {
	for i, line := range lines {
		err := scv.ComputeForLine(i+1, line)
		if err != nil {
			return err
		}
	}

	return nil
}

// ComputeForLine calculates the value for the line provided and returns the result.
// It also adds the value to the accumulator of the ScratchCardValidator.
func (scv *ScratchCardValidator) ComputeForLine(lineNumber int, line string) error {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return nil
	}

	// split game number from the cube sets
	result := strings.Split(line, ":")
	if len(result) != 2 {
		return fmt.Errorf("card line expected to contain a single ':' character")
	}

	cardNumberString := strings.TrimPrefix(result[0], "Card")
	cardNumberString = strings.TrimSpace(cardNumberString)
	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
		return fmt.Errorf("card number %q is not valid: %w", cardNumberString, err)
	}

	resultGameContentArr := strings.Split(result[1], "|")
	if len(resultGameContentArr) != 2 {
		return fmt.Errorf("card line expected to contain a single '|' character")
	}

	winningNumbers := make(map[int]struct{}) // set for fast lookup

	resultGameContentArr[0] = strings.TrimSpace(resultGameContentArr[0])
	resultWinningNumbersStringArr := strings.Split(resultGameContentArr[0], " ")
	for _, winningNumberString := range resultWinningNumbersStringArr {
		winningNumberString = strings.TrimSpace(winningNumberString)
		if winningNumberString == "" {
			continue
		}

		winningNumber, err := strconv.Atoi(winningNumberString)
		if err != nil {
			return fmt.Errorf("invalid winning number %q in card number '%d': %w", winningNumberString, cardNumber, err)
		}
		winningNumbers[winningNumber] = struct{}{}
	}

	var pickNumbers []int

	resultGameContentArr[1] = strings.TrimSpace(resultGameContentArr[1])
	resultPickNumbersStringArr := strings.Split(resultGameContentArr[1], " ")
	for _, pickNumberString := range resultPickNumbersStringArr {
		pickNumberString = strings.TrimSpace(pickNumberString)
		if pickNumberString == "" {
			continue
		}

		pickNumber, err := strconv.Atoi(pickNumberString)
		if err != nil {
			return fmt.Errorf("invalid pick number %q in card number '%d': %w", pickNumberString, cardNumber, err)
		}

		pickNumbers = append(pickNumbers, pickNumber)
	}

	if scv.debug {
		fmt.Printf("Card Number: %d -- Winning Numbers: %+v -- Pick Numbers: %+v\n", cardNumber, winningNumbers, pickNumbers)
	}

	// get points, add them and add to the accumulator
	points := 0

	firstMatch := true
	for _, number := range pickNumbers {
		if _, ok := winningNumbers[number]; ok {
			if firstMatch {
				points = 1
				firstMatch = false
			} else {
				points *= 2
			}
		}
	}

	scv.pointsAccumulator += points

	return nil
}

// PointsAccumulator returns the current result stored in the accumulator.
func (scv *ScratchCardValidator) PointsAccumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return scv.pointsAccumulator
}

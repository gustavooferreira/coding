package solutions

import (
	"strings"
	"unicode"
)

var digitConsts = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type Calibrator struct {
	accumulator int
}

func NewCalibrator() *Calibrator {
	return &Calibrator{}
}

// CalculateCalibrationArrayOfLines takes an array of lines and adds their calibration number to the accumulator.
func (c *Calibrator) CalculateCalibrationArrayOfLines(lines []string) {
	for _, line := range lines {
		c.CalculateCalibrationForLine(line)
	}
}

// CalculateCalibrationForLine calculates the calibration for the line provided and returns the result.
// It also adds the value to the accumulator of the calibrator.
func (c *Calibrator) CalculateCalibrationForLine(line string) int {
	firstDigit := true
	leftDigit := 0
	rightDigit := 0

	setDigit := func(d int) {
		if firstDigit {
			firstDigit = false
			leftDigit = d
			rightDigit = d
		} else {
			rightDigit = d
		}
	}

	for i, char := range line {
		if unicode.IsDigit(char) {
			d := int(char - '0')
			setDigit(d)
			continue
		}

		for j, digitSpelledOut := range digitConsts {
			if strings.HasPrefix(line[i:], digitSpelledOut) {
				setDigit(j + 1)
			}
		}
	}

	result := leftDigit*10 + rightDigit
	c.accumulator += result
	return result
}

// GetAccumulator returns the current result stored in the accumulator.
func (c *Calibrator) GetAccumulator() int {
	return c.accumulator
}

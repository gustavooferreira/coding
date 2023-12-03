package solutions

import (
	"unicode"
)

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

	for _, char := range line {
		if unicode.IsDigit(char) {
			d := int(char - '0')
			if firstDigit {
				firstDigit = false
				leftDigit = d
				rightDigit = d
			} else {
				rightDigit = d
			}
		}
	}

	result := leftDigit*10 + rightDigit
	c.accumulator += result
	return result
}

// GetAccumulator returns the current result stored in the accumulator.
func (c *Calibrator) GetAccumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return c.accumulator
}

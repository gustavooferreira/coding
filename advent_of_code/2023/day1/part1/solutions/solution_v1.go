package solutions

import (
	"strings"
	"unicode"
)

func CalculateCalibration(text string) int {
	textArr := strings.Split(text, "\n")

	result := 0

	for _, line := range textArr {
		result += NumberInLine(line)
	}

	return result
}

func NumberInLine(line string) int {
	var digits []int

	for _, char := range line {
		if unicode.IsDigit(char) {
			d := int(char - '0')
			digits = append(digits, d)
		}
	}

	if len(digits) == 0 {
		return 0
	} else if len(digits) == 1 {
		return digits[0]*10 + digits[0]
	}

	return digits[0]*10 + digits[len(digits)-1]
}

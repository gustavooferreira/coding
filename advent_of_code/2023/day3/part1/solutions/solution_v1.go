package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type PartNumberFinder struct {
	debug bool

	// internal representation of the schematic
	schematic  []string
	partNumber int
}

func NewPartNumberFinder() *PartNumberFinder {
	return &PartNumberFinder{}
}

func (pnf *PartNumberFinder) SetDebug(enable bool) {
	pnf.debug = enable
}

// LoadArrayOfLines loads the array of lines into the internal state.
func (pnf *PartNumberFinder) LoadArrayOfLines(lines []string) {
	for _, line := range lines {
		pnf.LoadLine(line)
	}
}

// LoadLine loads a line into the internal state.
func (pnf *PartNumberFinder) LoadLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	pnf.schematic = append(pnf.schematic, line)
}

// ComputePartNumber computes the part number.
func (pnf *PartNumberFinder) ComputePartNumber() error {
	// schematic is empty
	for len(pnf.schematic) == 0 {
		return nil
	}

	var numbers []number

	// loop over lines and find the numbers
	// build an array of structs with the numbers information
	for i, line := range pnf.schematic {
		numbersIndexes := findNumbersInLine(line)

		for _, nrIndx := range numbersIndexes {
			value, err := strconv.Atoi(line[nrIndx.Start : nrIndx.Stop+1])
			if err != nil {
				return fmt.Errorf("error converting string to number in line '%d': %w", i+1, err)
			}

			nr := number{
				Value:   value,
				Row:     i,
				Indexes: nrIndx,
			}

			numbers = append(numbers, nr)
		}
	}

	if pnf.debug {
		fmt.Println("All numbers found:")
		for _, number := range numbers {
			fmt.Printf("%+v\n", number)
		}
		fmt.Println("-----")
	}

	validNumbers := pnf.getValidNumbers(numbers)

	if pnf.debug {
		fmt.Println("All valid numbers:")
		for _, validNumber := range validNumbers {
			fmt.Printf("%+v\n", validNumber)
		}
	}

	for _, validNumber := range validNumbers {
		pnf.partNumber += validNumber
	}

	return nil
}

// getValidNumbers returns the numbers that belong to the part number.
func (pnf *PartNumberFinder) getValidNumbers(numbers []number) []int {
	var validNumbers []int

	// Compute limits
	cols := len(pnf.schematic[0])
	rows := len(pnf.schematic)

	for _, number := range numbers {
		start, stop := getNumberSurroundingLimits(number.Indexes, cols)

		// check line above number
		if number.Row != 0 {
			if foundSymbols(pnf.schematic[number.Row-1][start : stop+1]) {
				validNumbers = append(validNumbers, number.Value)
				continue
			}
		}

		// check line where the number is
		if foundSymbols(pnf.schematic[number.Row][start : stop+1]) {
			validNumbers = append(validNumbers, number.Value)
			continue
		}

		// check line below number
		if number.Row != rows-1 {
			if foundSymbols(pnf.schematic[number.Row+1][start : stop+1]) {
				validNumbers = append(validNumbers, number.Value)
				continue
			}
		}
	}

	return validNumbers
}

// PartNumber return the part number.
// Use this method after calling ComputePartNumber method.
func (pnf *PartNumberFinder) PartNumber() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return pnf.partNumber
}

type numberIndexes struct {
	Start int
	Stop  int
}

type number struct {
	// value represents the value of the number
	Value int
	// row represents the row the number is in
	Row int
	// indexes represent the line indexes for the number
	Indexes numberIndexes
}

// --- Helper functions ---

// Function to find the indexes of all numbers in the line.
func findNumbersInLine(line string) []numberIndexes {
	// state machine state
	currentlyProcessingNumber := false
	indexStart := 0
	indexStop := 0

	var numbersIndexes []numberIndexes

	for i, char := range line {
		if unicode.IsDigit(char) {
			if !currentlyProcessingNumber {
				indexStart = i
			}

			currentlyProcessingNumber = true
		} else {
			if currentlyProcessingNumber {
				indexStop = i - 1
				numbersIndexes = append(numbersIndexes, numberIndexes{
					Start: indexStart,
					Stop:  indexStop,
				})
			}

			currentlyProcessingNumber = false
		}
	}

	// if the number ends with the line, make sure to process that as well
	if currentlyProcessingNumber {
		indexStop = len(line) - 1
		numbersIndexes = append(numbersIndexes, numberIndexes{
			Start: indexStart,
			Stop:  indexStop,
		})
	}

	return numbersIndexes
}

// getNumberSurroundingLimits works out the limits of the number borders.
// This function is used to work out the range of characters to go over to detect symbols.
func getNumberSurroundingLimits(nrIdx numberIndexes, cols int) (start int, stop int) {
	start = nrIdx.Start - 1
	if start < 0 {
		start = 0
	}

	stop = nrIdx.Stop + 1
	if stop >= cols {
		stop = cols - 1
	}

	return start, stop
}

// foundSymbols reports whether it found symbols in the extracted substring.
func foundSymbols(substrline string) bool {
	for _, char := range substrline {
		if !unicode.IsDigit(char) && char != '.' {
			return true
		}
	}

	return false
}

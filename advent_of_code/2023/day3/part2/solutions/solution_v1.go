package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type GearRatioFinder struct {
	debug bool

	// internal representation of the schematic
	schematic    []string
	gearRatioSum int
}

func NewGearRatioFinder() *GearRatioFinder {
	return &GearRatioFinder{}
}

func (grf *GearRatioFinder) SetDebug(enable bool) {
	grf.debug = enable
}

// LoadArrayOfLines loads the array of lines into the internal state.
func (grf *GearRatioFinder) LoadArrayOfLines(lines []string) {
	for _, line := range lines {
		grf.LoadLine(line)
	}
}

// LoadLine loads a line into the internal state.
func (grf *GearRatioFinder) LoadLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	grf.schematic = append(grf.schematic, line)
}

// ComputeGearRatioSum computes the gear ratio sum.
func (grf *GearRatioFinder) ComputeGearRatioSum() error {
	// schematic is empty
	for len(grf.schematic) == 0 {
		return nil
	}

	var numbers []number

	// loop over lines and find the numbers
	// build an array of structs with the numbers information
	for i, line := range grf.schematic {
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

	if grf.debug {
		fmt.Println("All numbers found:")
		for _, number := range numbers {
			fmt.Printf("%+v\n", number)
		}
		fmt.Println("-----")
	}

	validNumbers := grf.getValidNumbers(numbers)

	if grf.debug {
		fmt.Println("All valid numbers:")
		for _, validNumber := range validNumbers {
			fmt.Printf("%+v\n", validNumber)
		}
		fmt.Println("-----")
	}

	validNumbersHash := make(map[string][]number)
	for _, validNumber := range validNumbers {
		// go over all gears for a number
		for _, gearCoordinate := range validNumber.GearsCoordinates {
			// compute key for * coordinates
			gearKey := fmt.Sprintf("%d,%d", gearCoordinate.y, gearCoordinate.x)

			validNumbersHash[gearKey] = append(validNumbersHash[gearKey], validNumber)
		}
	}

	if grf.debug {
		fmt.Println("All valid numbers in a hash:")
		for _, nrs := range validNumbersHash {
			fmt.Printf("%+v\n", nrs)
		}
	}

	// Go over the hashmap and if only 2 numbers found, multiply them
	for _, nrs := range validNumbersHash {
		if len(nrs) == 2 {
			grf.gearRatioSum += nrs[0].Value * nrs[1].Value
		}
	}

	return nil
}

// getValidNumbers returns the numbers that contain at least one gear around them.
func (grf *GearRatioFinder) getValidNumbers(numbers []number) []number {
	var validNumbers []number

	// Compute limits
	cols := len(grf.schematic[0])
	rows := len(grf.schematic)

	for _, number := range numbers {
		start, stop := getNumberSurroundingLimits(number.Indexes, cols)

		// check line above number
		if number.Row != 0 {
			indexes, found := foundSymbols(grf.schematic[number.Row-1][start:stop+1], start)
			if found {
				var gearsCoords []gearCoordinates
				for _, index := range indexes {
					gearsCoords = append(gearsCoords, gearCoordinates{
						x: index,
						y: number.Row - 1,
					})
				}

				number.GearsCoordinates = gearsCoords

				validNumbers = append(validNumbers, number)
				continue
			}
		}

		// check line where the number is
		indexes, found := foundSymbols(grf.schematic[number.Row][start:stop+1], start)
		if found {
			var gearsCoords []gearCoordinates
			for _, index := range indexes {
				gearsCoords = append(gearsCoords, gearCoordinates{
					x: index,
					y: number.Row,
				})
			}

			number.GearsCoordinates = gearsCoords

			validNumbers = append(validNumbers, number)
			continue
		}

		// check line below number
		if number.Row != rows-1 {
			indexes, found := foundSymbols(grf.schematic[number.Row+1][start:stop+1], start)
			if found {
				var gearsCoords []gearCoordinates
				for _, index := range indexes {
					gearsCoords = append(gearsCoords, gearCoordinates{
						x: index,
						y: number.Row + 1,
					})
				}

				number.GearsCoordinates = gearsCoords

				validNumbers = append(validNumbers, number)
				continue
			}
		}
	}

	return validNumbers
}

// GearRatioSum return the gear ratio sum.
// Use this method after calling GearRatioSum method.
func (grf *GearRatioFinder) GearRatioSum() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return grf.gearRatioSum
}

type gearCoordinates struct {
	x int
	y int
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
	// Contains all the gears that surround this number
	GearsCoordinates []gearCoordinates
}

// --- Helper functions ---

// Function to find the indexes of all numbers in the line.
// It only takes into account numbers that are surrounded by *
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

// foundSymbols reports whether it found at least one gear in the extracted substring.
// It also returns the gears index.
// offset represents the offset of the substrline from the beginning of the actual line
func foundSymbols(substrline string, offset int) (indexes []int, found bool) {
	for i, char := range substrline {
		if char == '*' {
			indexes = append(indexes, offset+i)
		}
	}

	return indexes, len(indexes) != 0
}

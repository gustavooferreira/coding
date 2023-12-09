package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	// each line represents a value changing over time
	puzzle [][]int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) SetDebug(enable bool) {
	s.debug = enable
}

// LoadArrayOfLines loads the array in the internal representation of the input content.
func (s *Solver) LoadArrayOfLines(lines []string) {
	for _, line := range lines {
		s.LoadLine(line)
	}
}

// LoadLine loads the input line into the internal representation of the input content.
func (s *Solver) LoadLine(line string) {
	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	// parse input
	itemIndex := 0

	for _, line := range s.inputContent {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		items := strings.Fields(line)

		var numbersLine []int

		for _, item := range items {
			number, _ := strconv.Atoi(item)
			numbersLine = append(numbersLine, number)
		}

		s.puzzle = append(s.puzzle, numbersLine)
		itemIndex++
	}

	if s.debug {
		fmt.Println("Input:")

		for _, line := range s.puzzle {
			fmt.Printf("%+v\n", line)
		}
		fmt.Println("----------")
	}

	result := 0

	// Work out prediction for each line
	for _, line := range s.puzzle {
		prediction := s.predictPreviousNumber(line)
		result += prediction
	}

	s.result = result
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

func (s *Solver) predictPreviousNumber(line []int) int {
	result := 0

	finished := false
	currentLine := line
	var intermediaryLines [][]int

	if s.debug {
		fmt.Printf("Computed prediction for: %+v\n", line)
	}

	for !finished {
		var resultLine []int
		diffThanZero := false

		for i := 0; i < len(currentLine)-1; i++ {
			diff := currentLine[i+1] - currentLine[i]
			resultLine = append(resultLine, diff)
			if diff != 0 {
				diffThanZero = true
			}
		}

		currentLine = resultLine
		intermediaryLines = append(intermediaryLines, resultLine)

		if s.debug {
			fmt.Printf("Iteration line result: %+v\n", resultLine)
		}

		// decide whether we are done
		if !diffThanZero {
			finished = true
		}
	}

	holdingValue := 0

	// compute prediction for previous number
	for i := len(intermediaryLines) - 2; i >= 0; i-- {
		holdingValue = intermediaryLines[i][0] - holdingValue
	}

	result = line[0] - holdingValue

	if s.debug {
		fmt.Printf("Prediction: %d\n", result)
	}

	return result
}

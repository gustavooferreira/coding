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

	lines []LineInfo
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
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	for _, line := range s.inputContent {
		lineArr := strings.Split(line, " ")

		unknownCellsCount := 0
		var layout []CellType

		for _, cell := range lineArr[0] {
			layout = append(layout, CellType(cell))

			if cell == '?' {
				unknownCellsCount++
			}
		}

		formatArr := strings.Split(lineArr[1], ",")

		var format []int
		for _, numberStr := range formatArr {
			number, _ := strconv.Atoi(numberStr)
			format = append(format, number)
		}

		s.lines = append(s.lines, LineInfo{
			Layout:            layout,
			Format:            format,
			UnknownCellsCount: unknownCellsCount,
		})
	}

	if s.debug {
		fmt.Println("Lines:")

		for i, lineInfo := range s.lines {
			fmt.Printf("%3d --- %+v\n", i, lineInfo)
		}
	}

	// TODO: solve the rest here

	// for each line, pass to a function that will generate all possible combinations and then runs each
	// combination through the validation function.

	// if valid, add the number of combinations to the result.

	s.result = 0
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

// ValidateSolvedLine validates whether the solved line is valid.
// That is, given the format, it chech that the layour conforms with that format.
func ValidateSolvedLine(layout []CellType, format []int) bool {

	return false
}

type LineInfo struct {
	Layout            []CellType
	Format            []int
	UnknownCellsCount int
}

type CellType string

const (
	CellTypeOperational CellType = "."
	CellTypeDamaged     CellType = "#"
	CellTypeUnknown     CellType = "?"
)

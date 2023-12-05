package solutions

import (
	"fmt"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int
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
	for i, line := range s.inputContent {
		s.ComputeLine(i+1, line)
	}
}

// ComputeLine computes the result for line.
// This method may not be used at times.
func (s *Solver) ComputeLine(lineNumber int, line string) int {
	result := 0

	if s.debug {
		fmt.Printf("Line: %3d -- Number: %d\n", lineNumber, result)
	}

	s.result += result
	return result
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

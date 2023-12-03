package solutions

import (
	"fmt"
)

type Solver struct {
	debug       bool
	accumulator int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (c *Solver) SetDebug(enable bool) {
	c.debug = enable
}

// LoadArrayOfLines takes an array of lines and adds their value number to the accumulator.
func (c *Solver) LoadArrayOfLines(lines []string) {
	for i, line := range lines {
		c.ComputeForLine(i+1, line)
	}
}

// ComputeForLine calculates the value for the line provided and returns the result.
// It also adds the value to the accumulator of the Solver.
func (c *Solver) ComputeForLine(lineNumber int, line string) int {
	result := 0

	if c.debug {
		fmt.Printf("Line: %3d -- Number: %d\n", lineNumber, result)
	}

	c.accumulator += result
	return result
}

// Accumulator returns the current result stored in the accumulator.
func (c *Solver) Accumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return c.accumulator
}

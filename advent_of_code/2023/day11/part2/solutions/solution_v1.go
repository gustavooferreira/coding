package solutions

import (
	"fmt"
	"math"
	"strings"
)

const (
	Spacing int = 1000000 - 1
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	// the index for a galaxy represents its number
	galaxies []Position

	// Represents a set. If the value is in the set, then the row is empty.
	emptyRows map[int]struct{}
	// Represents a set. If the value is in the set, then the col is empty.
	emptyCols map[int]struct{}
}

func NewSolver() *Solver {
	return &Solver{
		emptyRows: make(map[int]struct{}),
		emptyCols: make(map[int]struct{}),
	}
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
	// Find empty rows
	for j, line := range s.inputContent {
		emptyRow := true
		for _, cell := range line {
			if cell == '#' {
				emptyRow = false
			}
		}

		if emptyRow {
			s.emptyRows[j] = struct{}{}
		}
	}

	// Find empty cols
	for i := 0; i < len(s.inputContent[0]); i++ {
		emptyCol := true
		for j := 0; j < len(s.inputContent); j++ {
			cell := s.inputContent[j][i]
			if cell == '#' {
				emptyCol = false
			}
		}

		if emptyCol {
			s.emptyCols[i] = struct{}{}
		}
	}

	// find galaxies in the expanded universe
	expandedRows := 0

	for j, line := range s.inputContent {
		if _, ok := s.emptyRows[j]; ok {
			expandedRows += Spacing
			continue
		}

		expandedCols := 0

		for i, cell := range line {
			if _, ok := s.emptyCols[i]; ok {
				expandedCols += Spacing
				continue
			}

			if cell == '#' { // found a galaxy
				s.galaxies = append(s.galaxies, Position{
					X: i + expandedCols,
					Y: j + expandedRows,
				})
			}
		}
	}

	if s.debug {
		fmt.Println("Galaxies:")
		for i, galaxyPos := range s.galaxies {
			fmt.Printf("# %3d --- PosX: %3d , PosY: %3d\n", i, galaxyPos.X, galaxyPos.Y)
		}

		fmt.Println("Empty Rows:")
		for emptyRowIdx := range s.emptyRows {
			fmt.Printf("%-3d ", emptyRowIdx)
		}
		fmt.Println()

		fmt.Println("Empty Cols:")
		for emptyColIdx := range s.emptyCols {
			fmt.Printf("%-3d ", emptyColIdx)
		}
		fmt.Println()
	}

	// Create the permutations for galaxy pairs
	for i := 0; i < len(s.galaxies)-1; i++ {
		for j := i + 1; j < len(s.galaxies); j++ {
			galaxy1Pos := s.galaxies[i]
			galaxy2Pos := s.galaxies[j]

			distance := FindShortestPath(galaxy1Pos, galaxy2Pos)
			s.result += distance
		}
	}
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

type Position struct {
	X int
	Y int
}

// FindShortestPath finds the shortest path between 2 points.
func FindShortestPath(pos1 Position, pos2 Position) int {
	diffX := int(math.Abs(float64(pos1.X) - float64(pos2.X)))
	diffY := int(math.Abs(float64(pos1.Y) - float64(pos2.Y)))
	return diffX + diffY
}

package solutions

import (
	"fmt"
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	puzzles []Puzzle
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
	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	// split input into puzzles
	puzzle := Puzzle{}
	for _, line := range s.inputContent {
		if line == "" {
			s.puzzles = append(s.puzzles, puzzle)
			puzzle = Puzzle{}
			continue
		}

		puzzle.AddLine(line)
	}

	if puzzle.Length() > 0 {
		s.puzzles = append(s.puzzles, puzzle)
		puzzle = Puzzle{}
	}

	if s.debug {
		fmt.Println("Puzzles:")

		for i, puzzle := range s.puzzles {
			fmt.Printf("|- %d\n", i)
			fmt.Printf("%s", puzzle)
			fmt.Printf("|----\n\n")
		}
	}

	result := 0

	for _, puzzle := range s.puzzles {
		res := puzzle.ComputeResult()
		result += res
	}

	s.result = result
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

type Puzzle struct {
	contentRows    []string
	contentColumns []string
}

func (p *Puzzle) AddLine(line string) {
	p.contentRows = append(p.contentRows, line)

	if len(p.contentColumns) == 0 {
		p.contentColumns = make([]string, len(line))
	}

	for i, char := range line {
		p.contentColumns[i] = fmt.Sprintf("%s%s", p.contentColumns[i], string(char))
	}
}

func (p *Puzzle) Length() int {
	return len(p.contentRows)
}

func (p *Puzzle) CheckMirror(mirrorType MirrorType, firstIndex int) bool {
	indexPointerLeft := firstIndex
	indexPointerRight := firstIndex + 1

	if mirrorType == MirrorTypeHorizontal {
		for {
			indexPointerLeft--
			indexPointerRight++

			if indexPointerLeft < 0 || indexPointerRight >= len(p.contentRows) {
				return true
			}

			if p.contentRows[indexPointerLeft] != p.contentRows[indexPointerRight] {
				return false
			}
		}
	} else if mirrorType == MirrorTypeVertical {
		for {
			indexPointerLeft--
			indexPointerRight++

			if indexPointerLeft < 0 || indexPointerRight >= len(p.contentColumns) {
				return true
			}

			if p.contentColumns[indexPointerLeft] != p.contentColumns[indexPointerRight] {
				return false
			}
		}
	}

	// we should never hit this
	return true
}

func (p *Puzzle) ComputeResult() int {
	_, originalIdx, originalMirrorType := p.ComputeResultForCurrentState()

	// loop over every single cell
	for j, line := range p.contentRows {
		for i := range line {
			cell := p.contentRows[j][i]
			flippedCell := FlipCell(cell)

			// mutate cell for both rows and cols
			rowBytes := []byte(p.contentRows[j])
			rowBytes[i] = flippedCell
			p.contentRows[j] = string(rowBytes)

			colBytes := []byte(p.contentColumns[i])
			colBytes[j] = flippedCell
			p.contentColumns[i] = string(colBytes)
			// -----------

			// call compute result, if zero, it means we didn't find anything.
			res, idx, mirrorType := p.ComputeResultForCurrentState()
			if res != 0 && !(idx == originalIdx && mirrorType == originalMirrorType) {
				return res
			}

			// put cell back
			rowBytes[i] = cell
			p.contentRows[j] = string(rowBytes)
			colBytes[j] = cell
			p.contentColumns[i] = string(colBytes)
		}
	}

	return 0
}

func (p *Puzzle) ComputeResultForCurrentState() (result int, idx int, mirrorType MirrorType) {

	// Go through the rows.
	// If there is a mirror, get the index and multiplied by 100.
	for i, row := range p.contentRows {
		if i >= len(p.contentRows)-1 {
			continue
		}

		if row == p.contentRows[i+1] && p.CheckMirror(MirrorTypeHorizontal, i) {
			return 100 * (i + 1), i + 1, MirrorTypeHorizontal
		}
	}

	// Go through the cols.
	// If there is a mirror, get the index.
	for i, cols := range p.contentColumns {
		if i >= len(p.contentColumns)-1 {
			continue
		}

		if cols == p.contentColumns[i+1] && p.CheckMirror(MirrorTypeVertical, i) {
			return i + 1, i + 1, MirrorTypeVertical
		}
	}

	return 0, 0, ""
}

func (p Puzzle) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintln("Rows:"))
	for _, line := range p.contentRows {
		sb.WriteString(fmt.Sprintf("%s\n", line))
	}
	sb.WriteString(fmt.Sprintln())

	sb.WriteString(fmt.Sprintln("Cols:"))
	for _, col := range p.contentColumns {
		sb.WriteString(fmt.Sprintf("%s\n", col))
	}

	return sb.String()
}

type MirrorType string

const (
	MirrorTypeVertical   MirrorType = "mirror-vertical"
	MirrorTypeHorizontal MirrorType = "mirror-horizontal"
)

// FlipCell returns a . if it gets a #, or it returns a # if it gets a .
func FlipCell(cell byte) byte {
	if cell == '.' {
		return '#'
	} else if cell == '#' {
		return '.'
	}

	// this should never be hit
	return ' '
}

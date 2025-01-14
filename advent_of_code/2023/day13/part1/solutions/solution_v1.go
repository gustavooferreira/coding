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
	puzzle := Puzzle{debug: s.debug}
	for _, line := range s.inputContent {
		if line == "" {
			s.puzzles = append(s.puzzles, puzzle)
			puzzle = Puzzle{debug: s.debug}
			continue
		}

		puzzle.AddLine(line)
	}

	if puzzle.Length() > 0 {
		s.puzzles = append(s.puzzles, puzzle)
		puzzle = Puzzle{debug: s.debug}
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

	for i, puzzle := range s.puzzles {
		if s.debug {
			fmt.Printf("Computing for puzzle: %2d\n", i)
		}
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
	debug   bool
	content []string
}

func (p *Puzzle) AddLine(line string) {
	p.content = append(p.content, line)
}

func (p *Puzzle) Length() int {
	return len(p.content)
}

func (p *Puzzle) CheckMirror(mirrorType MirrorType, firstIndex int) bool {
	indexPointerLeft := firstIndex
	indexPointerRight := firstIndex + 1

	if mirrorType == MirrorTypeHorizontal {
		for {
			indexPointerLeft--
			indexPointerRight++

			if indexPointerLeft < 0 || indexPointerRight >= len(p.content) {
				return true
			}

			if p.content[indexPointerLeft] != p.content[indexPointerRight] {
				return false
			}
		}
	} else if mirrorType == MirrorTypeVertical {
		for {
			indexPointerLeft--
			indexPointerRight++

			if indexPointerLeft < 0 || indexPointerRight >= len(p.content[0]) {
				return true
			}

			line1 := []byte{}
			line2 := []byte{}

			// go through each column ==>> look right
			for j := range p.content {
				col1Cell := p.content[j][indexPointerLeft]
				col2Cell := p.content[j][indexPointerRight]

				line1 = append(line1, col1Cell)
				line2 = append(line2, col2Cell)
			}

			if string(line1) != string(line2) {
				return false
			}
		}
	}

	// we should never hit this
	return true
}

func (p *Puzzle) ComputeResult() int {
	result := 0

	// Go through the rows.
	// If there is a mirror, get the index and multiplied by 100.
	for i, row := range p.content {
		if i >= len(p.content)-1 {
			continue
		}

		if row == p.content[i+1] && p.CheckMirror(MirrorTypeHorizontal, i) {
			result += 100 * (i + 1)
			if p.debug {
				fmt.Printf("Row Mirror: %2d\n", i+1)
			}

			break
		}
	}

	// Go through the cols.
	// If there is a mirror, get the index.
	for i := range p.content[0] {
		if i >= len(p.content[0])-1 {
			continue
		}

		line1 := []byte{}
		line2 := []byte{}

		for j := range p.content {
			col1Cell := p.content[j][i]
			col2Cell := p.content[j][i+1]

			line1 = append(line1, col1Cell)
			line2 = append(line2, col2Cell)
		}

		if string(line1) == string(line2) && p.CheckMirror(MirrorTypeVertical, i) {
			result += i + 1
			if p.debug {
				fmt.Printf("Col Mirror: %2d\n", i+1)
			}
			break
		}
	}

	return result
}

func (p Puzzle) String() string {
	var sb strings.Builder

	for _, line := range p.content {
		sb.WriteString(fmt.Sprintf("%s\n", line))
	}

	return sb.String()
}

type MirrorType string

const (
	MirrorTypeVertical   MirrorType = "mirror-vertical"
	MirrorTypeHorizontal MirrorType = "mirror-horizontal"
)

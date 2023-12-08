package solutions

import (
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	directions string
	mappings   map[string]Node
}

func NewSolver() *Solver {
	return &Solver{
		mappings: make(map[string]Node),
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
	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	// Parse first line
	s.directions = strings.TrimSpace(s.inputContent[0])

	// Parse all other lines and put these in a map
	for i, line := range s.inputContent {
		if i == 0 {
			continue
		}

		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// split on =
		split1 := strings.Split(line, "=")
		key := strings.TrimSpace(split1[0])
		value := strings.TrimSpace(split1[1])
		value = strings.Replace(value, "(", "", -1)
		value = strings.Replace(value, ")", "", -1)

		// split on comma (trim parenthesis)
		valueArr := strings.Split(value, ",")
		valueL := strings.TrimSpace(valueArr[0])
		valueR := strings.TrimSpace(valueArr[1])

		s.mappings[key] = Node{
			Left:  valueL,
			Right: valueR,
		}
	}

	found := false
	steps := 1
	position := "AAA"

	for !found {
		for _, direction := range s.directions {
			node := s.mappings[position]

			if direction == 'R' {
				position = node.Right
			} else if direction == 'L' {
				position = node.Left
			}

			if position == "ZZZ" {
				found = true
				break
			}

			steps++
		}
	}

	s.result = steps
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

type Node struct {
	Left  string
	Right string
}

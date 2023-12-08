package solutions

import (
	"fmt"
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	directions string
	mappings   map[string]Node

	// This is initialised with the starting points positions
	positions         []string
	resultPerPosition []Result
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

		node := Node{
			Left:  valueL,
			Right: valueR,
		}

		s.mappings[key] = node

		// check if this is one of the starting point nodes
		if strings.HasSuffix(key, "A") {
			s.positions = append(s.positions, key)
		}
	}

	if s.debug {
		fmt.Println("Mappings:")
		for k, v := range s.mappings {
			fmt.Printf("Key: %5s -- Left: %5s -- Right: %5s\n", k, v.Left, v.Right)
		}

		fmt.Printf("Starting Positions: %+v\n", s.positions)
		fmt.Printf("Directions: %s\n", s.directions)
	}

	// work out how many steps per position
	for _, position := range s.positions {
		found := false
		steps := 1

		positionResult := Result{
			StartPosition: position,
		}

		for !found {
			for _, direction := range s.directions {
				if direction != 'R' && direction != 'L' {
					continue
				}

				node := s.mappings[position]

				if direction == 'R' {
					position = node.Right
				} else if direction == 'L' {
					position = node.Left
				}

				if strings.HasSuffix(position, "Z") {
					found = true
					break
				}

				steps++
			}
		}

		positionResult.StopPosition = position
		positionResult.Steps = steps

		s.resultPerPosition = append(s.resultPerPosition, positionResult)
	}

	if s.debug {
		fmt.Println("Results for each path:")
		for _, result := range s.resultPerPosition {
			fmt.Printf("Start: %4s -- Stop: %4s -- Steps: %6d\n", result.StartPosition, result.StopPosition, result.Steps)
		}
	}

	var results []int
	for _, r := range s.resultPerPosition {
		results = append(results, r.Steps)
	}

	s.result = LCM(results...)
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

type Result struct {
	StartPosition string
	StopPosition  string
	Steps         int
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) <= 1 {
		panic(fmt.Errorf("input must be at least 2 numbers"))
	}

	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

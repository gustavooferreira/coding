package solutions

import (
	"math"
	"strconv"
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	race Race
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
	// parse input into data structure
	timeString := strings.TrimSpace(strings.TrimPrefix(s.inputContent[0], "Time:"))
	distanceString := strings.TrimSpace(strings.TrimPrefix(s.inputContent[1], "Distance:"))

	// remove all spacing characters
	timeString = strings.Replace(timeString, " ", "", -1)
	distanceString = strings.Replace(distanceString, " ", "", -1)

	timeInt, _ := strconv.Atoi(timeString)
	distanceInt, _ := strconv.Atoi(distanceString)

	s.race = Race{
		Duration:     timeInt,
		bestDistance: distanceInt,
	}

	result := CountHowManyWaysOfWinning(s.race.Duration, s.race.bestDistance)

	s.result = result
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

type Race struct {
	Duration     int
	bestDistance int
}

func CountHowManyWaysOfWinning(raceDuration int, bestDistance int) int {
	var raceDurationF float64 = float64(raceDuration)
	var bestDistanceF float64 = float64(bestDistance)

	equationPart1 := raceDurationF / 2
	equationPart2 := math.Sqrt(math.Pow(raceDurationF, 2)-4*bestDistanceF) / 2

	root1 := equationPart1 - equationPart2
	root2 := equationPart1 + equationPart2

	score1 := int(math.Ceil(root1))
	score2 := int(math.Floor(root2))

	// we need to remove the edges if they don't beat the record but instead match it!
	if (raceDuration-score1)*score1 == bestDistance {
		score1 += 1
	}

	if (raceDuration-score2)*score2 == bestDistance {
		score2 -= 1
	}

	return score2 - score1 + 1
}

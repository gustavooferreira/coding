package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	// example: map seed-to-soil points to GenericMapper
	mappings map[string]*GenericMapper
	seeds    []ItemRange
}

func NewSolver() *Solver {
	mappings := make(map[string]*GenericMapper)
	return &Solver{
		mappings: mappings,
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

func (s *Solver) computeMappings() {
	// get seeds from first line
	seedsString := strings.TrimPrefix(s.inputContent[0], "seeds:")
	var seedsNumbers []int
	for _, seedString := range strings.Split(seedsString, " ") {
		seedString = strings.TrimSpace(seedString)
		if seedString == "" {
			continue
		}

		seedNumber, _ := strconv.Atoi(seedString)
		seedsNumbers = append(seedsNumbers, seedNumber)
	}

	// compute seed ranges
	for i := 0; i < len(seedsNumbers); i += 2 {
		s.seeds = append(s.seeds, ItemRange{
			Start:  seedsNumbers[i],
			Length: seedsNumbers[i+1],
		})
	}

	if s.debug {
		fmt.Printf("Seeds: %+v\n", s.seeds)
	}

	// go over the other lines and create the mappings
	currentMapping := ""

	for _, inputLine := range s.inputContent[1:] {
		inputLine = strings.TrimSpace(inputLine)
		if inputLine == "" {
			continue
		}

		if strings.HasSuffix(inputLine, "map:") {
			inputLine = strings.TrimSuffix(inputLine, "map:")
			inputLine = strings.TrimSpace(inputLine)
			currentMapping = inputLine
			continue
		}

		// we want to process the numbersString
		var numbers []int // we will assume we always get 3 numbers

		numbersString := strings.Split(inputLine, " ")
		for _, numberString := range numbersString {
			numberString = strings.TrimSpace(numberString)
			if numberString == "" {
				continue
			}

			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		// add mapping
		currentMappingArr := strings.Split(currentMapping, "-to-")
		source := strings.TrimSpace(currentMappingArr[0])
		destination := strings.TrimSpace(currentMappingArr[1])
		if _, ok := s.mappings[source]; !ok {
			s.mappings[source] = &GenericMapper{
				Source:      source,
				Destination: destination,
			}
		}

		s.mappings[source].AddRangeMapping(numbers[0], numbers[1], numbers[2])
	}
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	s.computeMappings()

	// // compute the "linked list result" and store the seed number and the location together.
	// var seedsLocationPair []SeedLocation
	//
	// for _, seedNumber := range seedsNumbers {
	// 	// walk through the mappings and get the location
	// 	source := "seed"
	// 	number := seedNumber
	//
	// 	for {
	// 		mapper := s.mappings[source]
	// 		number = mapper.Lookup(number)
	// 		source = mapper.Destination
	//
	// 		if mapper.Destination == "location" {
	// 			// store result
	// 			seedsLocationPair = append(seedsLocationPair, SeedLocation{
	// 				Seed:     seedNumber,
	// 				Location: number,
	// 			})
	//
	// 			break
	// 		}
	// 	}
	// }
	//
	// // find the minimum location and return the seed number
	// minLocation := -1
	//
	// for _, seedsLocationPair := range seedsLocationPair {
	// 	if minLocation == -1 || seedsLocationPair.Location < minLocation {
	// 		minLocation = seedsLocationPair.Location
	// 	}
	// }
	//
	// s.result = minLocation
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

type ItemRange struct {
	Start  int
	Length int
}

// GetRange returns the start and end of the range (start and end are inclusive).
func (sr *ItemRange) GetRange() (start int, end int) {
	return sr.Start, sr.Start + sr.Length - 1
}

type SeedLocation struct {
	Seed     int
	Location int
}

// RangeMapperInfo represents a range and the corresponding mapped numbers.
type RangeMapperInfo struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

type GenericMapper struct {
	Source      string
	Destination string

	// ranges represents an ordered list of ranges.
	ranges []RangeMapperInfo
}

// AddRange adds a range to the generic mapper.
func (gm *GenericMapper) AddRangeMapping(dstStart int, srcStart int, length int) {
	rmi := RangeMapperInfo{
		DestinationStart: dstStart,
		SourceStart:      srcStart,
		Length:           length,
	}

	gm.ranges = append(gm.ranges, rmi)

	// Sort array (a bit wasteful to do it on every insertion, but oh well)
	sort.Slice(gm.ranges, func(i int, j int) bool {
		return gm.ranges[i].SourceStart < gm.ranges[j].SourceStart
	})
}

// Lookup performs a mapping lookup.
// Returns ranges of output.
func (gm *GenericMapper) Lookup(input ItemRange) (output []ItemRange) {

	// start with the first element and go over the various ranges

	// return the various ranges for different mappings

	// // look up based on ranges
	// for _, rng := range gm.ranges {
	// 	if input >= rng.SourceStart && input <= rng.SourceStart+rng.Length-1 {
	// 		return rng.DestinationStart + input - rng.SourceStart
	// 	}
	// }
	//
	// // if we couldn't find it then return the same input number
	// return []ItemRange{{
	// 	Start:  0,
	// 	Length: 0,
	// }}
	return nil
}

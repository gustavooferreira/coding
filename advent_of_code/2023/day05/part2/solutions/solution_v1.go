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

	// end up with a bunch of locations
	var locations []ItemRange

	source := "seed"
	items := s.seeds

	// go over the "linked list"
	for {
		mapper := s.mappings[source]

		var outputItems []ItemRange
		for _, itemRange := range items {
			subOutputItems := mapper.Lookup(itemRange)
			outputItems = append(outputItems, subOutputItems...)
		}

		source = mapper.Destination
		items = outputItems

		if mapper.Destination == "location" {
			// store result
			locations = append(locations, items...)
			break
		}
	}

	// don't need this but it's nice for the output
	sort.Slice(locations, func(i int, j int) bool {
		return locations[i].Start < locations[j].Start
	})

	if s.debug {
		fmt.Println("Locations:")
		for _, location := range locations {
			fmt.Printf("%+v\n", location)
		}
	}

	// find the minimum location and return the seed number
	minLocation := -1

	for _, location := range locations {
		if minLocation == -1 || location.Start < minLocation {
			minLocation = location.Start
		}
	}

	s.result = minLocation
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

type SeedsRangeLocation struct {
	Seed      ItemRange
	Locations []ItemRange
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
	rangeStart := input.Start
	rangeEnd := input.Start + input.Length - 1

	// edge case (starting at zero)
	if rangeStart < gm.ranges[0].SourceStart {
		output = append(output, ItemRange{
			Start:  rangeStart,
			Length: gm.ranges[0].SourceStart - rangeStart,
		})
	}

	// edge case (towards infinity)
	lastEdge := gm.ranges[len(gm.ranges)-1].SourceStart + gm.ranges[len(gm.ranges)-1].Length // pointer to first number
	if rangeEnd >= lastEdge {
		output = append(output, ItemRange{
			Start:  lastEdge,
			Length: rangeEnd - lastEdge + 1,
		})
	}

	// compute overlap inside boxes
	for _, rng := range gm.ranges {
		boxStart := rng.SourceStart
		boxEnd := rng.SourceStart + rng.Length - 1

		item, valid := computeOverlap(rangeStart, rangeEnd, boxStart, boxEnd)
		if !valid { // range falls outside the box
			continue
		}

		// shift start index
		item.Start = item.Start + (rng.DestinationStart - rng.SourceStart)

		output = append(output, item)
	}

	// compute space between boxes
	for i, rng := range gm.ranges {
		// skip computing last segment in-between boxes (we took care of that above in the edge case)
		if i >= len(gm.ranges)-1 {
			continue
		}

		boxStart := rng.SourceStart + rng.Length
		boxEnd := gm.ranges[i+1].SourceStart - 1

		item, valid := computeOverlap(rangeStart, rangeEnd, boxStart, boxEnd)
		if !valid { // range falls outside the box
			continue
		}

		output = append(output, item)
	}

	return output
}

func computeOverlap(rangeStart int, rangeEnd int, boxStart int, boxEnd int) (itemRange ItemRange, valid bool) {
	// range falls outside the box
	if rangeEnd < boxStart || rangeStart > boxEnd {
		return ItemRange{}, false
	}

	start := 0
	end := 0

	if rangeStart < boxStart {
		start = boxStart
	} else {
		start = rangeStart
	}

	if rangeEnd > boxEnd {
		end = boxEnd
	} else {
		end = rangeEnd
	}

	return ItemRange{
		Start:  start,
		Length: end - start + 1,
	}, true
}

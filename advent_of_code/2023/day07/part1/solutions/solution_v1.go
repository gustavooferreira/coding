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

	// sorted (by strength, first entry is less strong than last) slice of hands
	hands []HandInfo
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
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		s.LoadLine(line)
	}
}

// LoadLine loads the input line into the internal representation of the input content.
func (s *Solver) LoadLine(line string) {
	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	// parse lines and put them in a struct
	for _, line := range s.inputContent {
		line = strings.TrimSpace(line)
		lineArr := strings.Fields(line)
		bid, _ := strconv.Atoi(lineArr[1])

		s.hands = append(s.hands, HandInfo{
			HandRepresentation: lineArr[0],
			Bid:                bid,
			Type:               0,
		})
	}

	// compute the strengths of the hand
	for i, hand := range s.hands {
		handType := computeCardHandStrength(hand.HandRepresentation)
		s.hands[i].Type = handType
	}

	// sort hands slice based on described logic
	sort.Slice(s.hands, func(i, j int) bool {
		hand1 := s.hands[i]
		hand2 := s.hands[j]

		if hand1.Type == hand2.Type {
			// go over each card in a hand
			for i, hand1CardC := range hand1.HandRepresentation {
				hand1Card := CardType(hand1CardC)
				hand2Card := CardType(rune(hand2.HandRepresentation[i]))

				if hand1Card < hand2Card {
					return true
				} else if hand1Card > hand2Card {
					return false
				}
			}

		} else if hand1.Type < hand2.Type {
			return true
		}

		return false
	})

	if s.debug {
		fmt.Println("Hands:")
		for _, hand := range s.hands {
			fmt.Printf("%+v\n", hand)
		}
	}

	// add up all the points
	s.result = 0
	for i, hand := range s.hands {
		s.result += hand.Bid * (i + 1)
	}
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

// create a map, key is the card type and value is the count

func computeCardHandStrength(cardHand string) HandType {
	cardTypeCount := make(map[rune]int)

	for _, cardType := range cardHand {
		cardTypeCount[cardType] += 1
	}

	foundThreeCount := 0
	foundTwoCount := 0

	for _, count := range cardTypeCount {
		if count == 5 {
			return HandType_FiveOfAKind
		} else if count == 4 {
			return HandType_FourOfAKind
		} else if count == 3 {
			foundThreeCount += 1
		} else if count == 2 {
			foundTwoCount += 1
		}
	}

	if foundThreeCount == 1 && foundTwoCount == 1 {
		return HandType_FullHouse
	} else if foundThreeCount == 1 {
		return HandType_ThreeOfAKind
	} else if foundTwoCount == 2 {
		return HandType_TwoPair
	} else if foundTwoCount == 1 {
		return HandType_OnePair
	}

	return HandType_HighCard
}

type HandInfo struct {
	HandRepresentation string
	Bid                int
	Type               HandType
}

type HandType int

const (
	HandType_HighCard HandType = iota
	HandType_OnePair
	HandType_TwoPair
	HandType_ThreeOfAKind
	HandType_FullHouse
	HandType_FourOfAKind
	HandType_FiveOfAKind
)

func (ht HandType) String() string {
	return [...]string{"High card", "One pair", "Two pair", "Three of a kind", "Full house", "Four of a kind", "Five of a kind"}[ht]
}

type cardType int

const (
	CardType_2 cardType = iota
	CardType_3
	CardType_4
	CardType_5
	CardType_6
	CardType_7
	CardType_8
	CardType_9
	CardType_T
	CardType_J
	CardType_Q
	CardType_K
	CardType_A
)

func CardType(card rune) cardType {
	switch card {
	case '2':
		return CardType_2
	case '3':
		return CardType_3
	case '4':
		return CardType_4
	case '5':
		return CardType_5
	case '6':
		return CardType_6
	case '7':
		return CardType_7
	case '8':
		return CardType_8
	case '9':
		return CardType_9
	case 'T':
		return CardType_T
	case 'J':
		return CardType_J
	case 'Q':
		return CardType_Q
	case 'K':
		return CardType_K
	case 'A':
		return CardType_A
	default:
		return -1
	}
}

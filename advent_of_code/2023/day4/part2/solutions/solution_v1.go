package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type ScratchCardCounter struct {
	debug bool

	// one entry per line of input
	scratchCards []ScratchCard

	scratchcardAccumulator int
}

type ScratchCard struct {
	// CardNumber represents the card number (ID)
	CardNumber int
	// Wins represents the number of winning numbers
	Wins int
	// Count represents the count of this scratch card.
	Count int
}

func NewScratchCardCounter() *ScratchCardCounter {
	return &ScratchCardCounter{}
}

func (scc *ScratchCardCounter) SetDebug(enable bool) {
	scc.debug = enable
}

// LoadArrayOfLines takes an array of lines and adds their value number to the accumulator.
func (scc *ScratchCardCounter) LoadArrayOfLines(lines []string) error {
	for i, line := range lines {
		err := scc.ComputeForLine(i+1, line)
		if err != nil {
			return err
		}
	}

	return nil
}

// ComputeForLine calculates the value for the line provided and returns the result.
// It also adds the value to the accumulator of the Solver.
func (scc *ScratchCardCounter) ComputeForLine(lineNumber int, line string) error {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return nil
	}

	cardNumber, winningNumbersCount, err := scc.computeNumberOfWinningNumbers(line)
	if err != nil {
		return fmt.Errorf("error while computing number of winning numbers: %w", err)
	}

	scc.scratchCards = append(scc.scratchCards, ScratchCard{
		CardNumber: cardNumber,
		Wins:       winningNumbersCount,
		Count:      1,
	})

	return nil
}

// ComputeFinalResult computes the final number of scratch cards.
func (scc *ScratchCardCounter) ComputeFinalResult() {
	if scc.debug {
		fmt.Println("Scratch Cards Result:")
		for _, scratchCard := range scc.scratchCards {
			fmt.Printf("Card Number: %d -- Winning Number Count: %d\n", scratchCard.CardNumber, scratchCard.Wins)
		}
	}

	// go over each line and update count
	for i, scratchCard := range scc.scratchCards {
		if scratchCard.Wins > 0 {
			// go over the below lines and increments count multiplied but its current count
			for j := 0; j < scratchCard.Wins; j++ {
				scc.scratchCards[i+j+1].Count += scratchCard.Count
			}
		}
	}

	if scc.debug {
		fmt.Println("Scratch Cards Final Result:")
		for _, scratchCard := range scc.scratchCards {
			fmt.Printf("Card Number: %3d -- Winning Number Count: %3d -- Card Count: %3d\n",
				scratchCard.CardNumber, scratchCard.Wins, scratchCard.Count)
		}
	}

	// go over one last time and add all the counts!
	result := 0
	for _, scratchCard := range scc.scratchCards {
		result += scratchCard.Count
	}

	scc.scratchcardAccumulator = result
}

// ScratchCardAccumulator returns the current result stored in the accumulator.
func (scc *ScratchCardCounter) ScratchCardAccumulator() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return scc.scratchcardAccumulator
}

func (scc *ScratchCardCounter) computeNumberOfWinningNumbers(line string) (int, int, error) {
	// split game number from the cube sets
	result := strings.Split(line, ":")
	if len(result) != 2 {
		return 0, 0, fmt.Errorf("card line expected to contain a single ':' character")
	}

	cardNumberString := strings.TrimPrefix(result[0], "Card")
	cardNumberString = strings.TrimSpace(cardNumberString)
	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
		return 0, 0, fmt.Errorf("card number %q is not valid: %w", cardNumberString, err)
	}

	resultGameContentArr := strings.Split(result[1], "|")
	if len(resultGameContentArr) != 2 {
		return 0, 0, fmt.Errorf("card line expected to contain a single '|' character")
	}

	winningNumbers := make(map[int]struct{}) // set for fast lookup

	resultGameContentArr[0] = strings.TrimSpace(resultGameContentArr[0])
	resultWinningNumbersStringArr := strings.Split(resultGameContentArr[0], " ")
	for _, winningNumberString := range resultWinningNumbersStringArr {
		winningNumberString = strings.TrimSpace(winningNumberString)
		if winningNumberString == "" {
			continue
		}

		winningNumber, err := strconv.Atoi(winningNumberString)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid winning number %q in card number '%d': %w", winningNumberString, cardNumber, err)
		}
		winningNumbers[winningNumber] = struct{}{}
	}

	var pickNumbers []int

	resultGameContentArr[1] = strings.TrimSpace(resultGameContentArr[1])
	resultPickNumbersStringArr := strings.Split(resultGameContentArr[1], " ")
	for _, pickNumberString := range resultPickNumbersStringArr {
		pickNumberString = strings.TrimSpace(pickNumberString)
		if pickNumberString == "" {
			continue
		}

		pickNumber, err := strconv.Atoi(pickNumberString)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid pick number %q in card number '%d': %w", pickNumberString, cardNumber, err)
		}

		pickNumbers = append(pickNumbers, pickNumber)
	}

	points := 0

	for _, number := range pickNumbers {
		if _, ok := winningNumbers[number]; ok {
			points += 1
		}
	}

	if scc.debug {
		fmt.Printf("Card Number: %d -- Winning Coung: %d -- Winning Numbers: %+v -- Pick Numbers: %+v\n", cardNumber, points, winningNumbers, pickNumbers)
	}

	return cardNumber, points, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day2/part2/solutions"
)

// Run example:
// > cat inputs/input2.txt | go run cmd/main.go
func main() {
	powerCalculator := solutions.NewPowerCalculator()

	lineNumber := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := powerCalculator.ComputeMinimumGameSetForLine(line)
		if err != nil {
			fmt.Printf("Error: error while checking game line number %d: %s\n", lineNumber, err)
			os.Exit(1)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", powerCalculator.GetGameSetPowerAccumulator())
}

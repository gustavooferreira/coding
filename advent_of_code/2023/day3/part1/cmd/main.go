package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day3/part1/solutions"
)

// Run example:
// > cat inputs/input2.txt | go run cmd/main.go
func main() {
	partNumberFinder := solutions.NewPartNumberFinder()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		partNumberFinder.LoadLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s\n", err)
		os.Exit(1)
	}

	err := partNumberFinder.ComputePartNumber()
	if err != nil {
		fmt.Printf("Error: error while computing part number: %s\n", err)
	}

	fmt.Printf("Result: %d\n", partNumberFinder.GetPartNumber())
}

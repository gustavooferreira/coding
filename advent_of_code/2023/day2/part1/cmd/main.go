package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day2/part1/solutions"
)

// Run example:
// > cat inputs/input2.txt | go run cmd/main.go
func main() {
	validator := solutions.NewValidator()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		validator.ValidateGameForLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", validator.GetGameIDAccumulator())
}

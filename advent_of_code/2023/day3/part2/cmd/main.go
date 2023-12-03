package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day3/part2/solutions"
)

// Run example:
// > cat inputs/input2.txt | go run cmd/main.go
func main() {
	gearRatioFinder := solutions.NewGearRatioFinder()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		gearRatioFinder.LoadLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s\n", err)
		os.Exit(1)
	}

	err := gearRatioFinder.ComputeGearRatioSum()
	if err != nil {
		fmt.Printf("Error: error while computing gear ratio sum: %s\n", err)
	}

	fmt.Printf("Result: %d\n", gearRatioFinder.GetGearRatioSum())
}

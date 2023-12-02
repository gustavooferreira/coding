package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day1/part2/solutions"
)

// Run example:
// > cat inputs/input2.txt | go run cmd/main.go
func main() {
	calibrator := solutions.NewCalibrator()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		calibrator.CalculateCalibrationForLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", calibrator.GetAccumulator())
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day04/part1/solutions"
)

// Run example:
// > cat solutions/testdata/input2.txt | go run cmd/main.go
func main() {
	debugFlag := flag.Bool("debug", false, "print debug messages")
	flag.Parse()

	scratchCardValidator := solutions.NewScratchCardValidator()
	scratchCardValidator.SetDebug(*debugFlag)

	lineNumber := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := scratchCardValidator.ComputeForLine(lineNumber, line)
		if err != nil {
			fmt.Printf("Error: error while processing line '%d': %s\n", lineNumber, err)
			os.Exit(1)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: error while scanning the file: %s\n", err)
		os.Exit(1)
	}

	if *debugFlag {
		fmt.Println("--------------------")
	}

	fmt.Printf("Result: %d\n", scratchCardValidator.PointsAccumulator())
}

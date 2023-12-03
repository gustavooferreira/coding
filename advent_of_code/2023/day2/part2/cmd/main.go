package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day2/part2/solutions"
)

// Run example:
// > cat solutions/testdata/input2.txt | go run cmd/main.go
func main() {
	debugFlag := flag.Bool("debug", false, "print debug messages")
	flag.Parse()

	powerCalculator := solutions.NewPowerCalculator()
	powerCalculator.SetDebug(*debugFlag)

	lineNumber := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := powerCalculator.ComputeMinimumGameSetForLine(lineNumber, line)
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

	if *debugFlag {
		fmt.Println("--------------------")
	}

	fmt.Printf("Result: %d\n", powerCalculator.GameSetPowerAccumulator())
}

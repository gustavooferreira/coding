package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/gustavooferreira/coding/advent_of_code/2023/day03/part2/solutions"
)

// Run example:
// > cat solutions/testdata/input2.txt | go run cmd/main.go
func main() {
	debugFlag := flag.Bool("debug", false, "print debug messages")
	flag.Parse()

	gearRatioFinder := solutions.NewGearRatioFinder()
	gearRatioFinder.SetDebug(*debugFlag)

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

	if *debugFlag {
		fmt.Println("--------------------")
	}

	fmt.Printf("Result: %d\n", gearRatioFinder.GearRatioSum())
}

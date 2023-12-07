package solutions

type NumberIndexes = numberIndexes

type Number = number

func FindNumbersInLine(line string) []numberIndexes {
	return findNumbersInLine(line)
}

func GetNumberSurroundingLimits(nrIdx numberIndexes, cols int) (start int, stop int) {
	return getNumberSurroundingLimits(nrIdx, cols)
}

func FoundSymbols(substrline string) bool {
	return foundSymbols(substrline)
}

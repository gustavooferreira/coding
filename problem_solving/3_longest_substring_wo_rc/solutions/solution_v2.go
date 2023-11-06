package solutions

func LengthOfLongestSubstring_V2(s string) int {
	maxLength := 0

	// We store all 128 code points for the ASCII set, which includes the letters, digits, symbols and spaces.
	// We could make an optimization here to leverage the fact that the map is initialized with all zeroes.
	// We could just increase all indices by one, which means if the entry is zero, it means the character hasn't
	// showned yet.
	charIndex := make([]int, 128)
	// initialize all entries to -1 as no indice can be less than zero. We will rely on this behaviour to assert
	// on whether a given character has already showed up in the string or not.
	for i := range charIndex {
		charIndex[i] = -1
	}

	leftPointer := 0

	for rightPointer, character := range s {
		if charIndex[character] >= leftPointer {
			// skip over the repeated character
			leftPointer = charIndex[character] + 1
		}

		charIndex[character] = rightPointer
		maxLength = max(maxLength, rightPointer-leftPointer+1)
	}

	return maxLength
}

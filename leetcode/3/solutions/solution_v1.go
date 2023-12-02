package solutions

func LengthOfLongestSubstring_V1(s string) int {
	countMax := 0

	for i := 0; i < len(s); i++ {
		str := s[i:]
		// count non repeating characters starting at the beginning of the string
		result := CountSequenceOfNonRepeatingCharacters(str)
		if result > countMax {
			countMax = result
		}
	}

	return countMax
}

func CountSequenceOfNonRepeatingCharacters(s string) int {
	hashmap := make(map[rune]struct{})
	total := 0

	for _, c := range s {
		_, ok := hashmap[c]
		if ok {
			return total
		}

		total += 1
		hashmap[c] = struct{}{}
	}

	return total
}

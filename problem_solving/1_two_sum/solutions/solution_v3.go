package solutions

// TwoSum_V3 puts all elements in a hashmap and then does a single pass while checking if the complement exists.
// This solution works because the problem statement says that it is guaranteed to only ever exist one solution,
// that means that at most we will have 2 numbers that are the same, otherwise we would end up with more than one
// solution. Using a hash map is fine, since we will start iterating from the beginning, which means the later values
// will be the ones in the hash map in the case of duplicates. And because on the second pass we start from the
// beginning again, it means we are always guaranteed to match against the second duplicate.
// This solution uses 2 passes, but we can further optimize this doing just one pass by checking for the complement
// at the same time that we are populating the hashmap.
func TwoSum_V3(nums []int, target int) []int {
	// key is the value and the value of the hashmap is the indice
	vals := make(map[int]int, len(nums))

	for i, v := range nums {
		vals[v] = i
	}

	// Do one pass on the nums slice
	for i, v := range nums {
		complement := target - v

		idx, ok := vals[complement]
		// if we don't have the complement, move on
		if !ok {
			continue
		}

		// check we aren't getting the value for the same position
		if idx == i {
			continue
		}

		return []int{i, idx}
	}

	// if we got here, it's because there was no solution
	return nil
}

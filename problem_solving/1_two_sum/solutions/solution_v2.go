package solutions

import (
	"sort"
)

// TwoSum sorts the nums array and then uses pointers on each end until we find a solution or the pointers converge.
func TwoSum_V2(nums []int, target int) []int {
	// Check that we get at least 2 elements in the nums slice
	if len(nums) < 2 {
		return nil
	}

	// We need to keep a struct that contains the old indices order because we will be sorting the slice.
	// Declare the struct here so the function is self-contained.
	type numRepresentation struct {
		val       int
		oldIndice int
	}

	var vals []numRepresentation

	for i, v := range nums {
		vals = append(vals, numRepresentation{
			val:       v,
			oldIndice: i,
		})
	}

	// Sort vals
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].val < vals[j].val
	})

	// Two pointers, starting at each end and start converging until we find the right solution?
	leftPointer := 0
	rightPointer := len(vals) - 1

	for leftPointer != rightPointer {
		// add the values in those indices
		result := vals[leftPointer].val + vals[rightPointer].val

		if result < target {
			leftPointer++
		} else if result > target {
			rightPointer--
		} else { // this is our solution
			return []int{vals[leftPointer].oldIndice, vals[rightPointer].oldIndice}
		}
	}

	// if we got here, it's because there was no solution
	return nil
}

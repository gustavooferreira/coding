package solutions

// TwoSum_V1 finds the result by bruteforce.
func TwoSum_V1(nums []int, target int) []int {
	// Check that we get at least 2 elements in the nums slice
	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// if we got here, it's because there was no solution
	return nil
}

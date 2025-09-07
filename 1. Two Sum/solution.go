func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for leftIdx := range len(nums) {
		left := nums[leftIdx]
		diff := target - left

		if rightIdx, ok := numsMap[diff]; ok {
			return []int{leftIdx, rightIdx}
		}

		numsMap[left] = leftIdx
	}

	return nil
}

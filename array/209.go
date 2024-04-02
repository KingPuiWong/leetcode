package array

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	i := 0
	result := len(nums) + 1
	subLength := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLength = j - i + 1
			if subLength < result {
				result = subLength
			}
			if i < len(nums) {
				sum -= nums[i]
				i++

			}
		}
	}

	if result == len(nums)+1 {
		return 0
	}

	return result
}

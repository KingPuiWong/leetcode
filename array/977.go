package array

func sortedSquares(nums []int) []int {
	low := 0
	high := len(nums) - 1
	k := len(nums) - 1
	result := make([]int, len(nums))
	for {
		if low < high {
			if nums[low]*nums[low] > nums[high]*nums[high] {
				result[k] = nums[low] * nums[low]
				low++
				k--
			} else {
				result[k] = nums[high] * nums[high]
				high--
				k--
			}
		} else if low == high {
			result[k] = nums[low] * nums[low]
			break
		}
	}

	return result

}

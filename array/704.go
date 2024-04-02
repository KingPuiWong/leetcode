package array

func search(nums []int, target int) int {
	lowIndex := 0
	highIndex := len(nums) - 1

	for lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2
		if nums[midIndex] < target {
			lowIndex = midIndex + 1
		} else if nums[midIndex] > target {
			highIndex = midIndex - 1
		} else if nums[midIndex] == target {
			return midIndex
		}
	}

	return -1
}

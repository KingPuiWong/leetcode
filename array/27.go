package array

// 双指针法
// 快指针: 寻找新数组的元素，新数组就是不含有目标元素的数组
// 慢指针: 指向更新 新数组下标的位置
func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	return slowIndex
}

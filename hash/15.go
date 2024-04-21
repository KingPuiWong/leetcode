package hash

import (
	"sort"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]

满足 i != j、i != k 且 j != k

同时还满足 nums[i] + nums[j] + nums[k] == 0

请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组
示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

!!! 这道题不适合用哈希,使用双指针法
*/
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	var res [][]int
	sort.Ints(nums)
	// 这里因为left,right;所以因为-2
	for i := 0; i < length-2; i++ {
		n1 := nums[i]
		// 因为排序后,如果第一个大于0,不可能组成三元组
		if n1 > 0 {
			break
		}

		//去重 a
		if i > 0 && nums[i-1] == n1 {
			continue
		}

		left := i + 1
		right := length - 1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 > 0 {
				right--
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				res = append(res, []int{n1, n2, n3})
				// 去重 b 和 c
				for left < right && nums[left] == n2 {
					left++
				}

				for left < right && nums[right] == n3 {
					right--
				}
			}
		}
	}

	return res
}

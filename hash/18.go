package hash

import (
	"sort"
)

/*
*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
（若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		n1 := nums[i]

		for k := i + 1; k < len(nums)-2; k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}

			n2 := nums[k]
			l := k + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l+1] {
						l++
					}

					for l < r && n4 == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}

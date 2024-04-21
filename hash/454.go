package hash

/**
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，
数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countAB := make(map[int]int)
	for _, v := range nums1 {
		for _, w := range nums2 {
			countAB[v+w]++
		}
	}

	var res int
	for _, v := range nums3 {
		for _, w := range nums4 {
			res += countAB[-(v + w)]
		}
	}

	return res
}

package hash

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	existMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		existMap[nums1[i]] = 1
	}

	resMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		if _, exist := existMap[nums2[i]]; exist {
			resMap[nums2[i]] = nums2[i]
		}
	}

	for _, value := range resMap {
		res = append(res, value)
	}

	return res
}

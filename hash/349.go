package hash

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	existMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, exist := existMap[nums1[i]]; !exist {
			existMap[nums1[i]] = 1
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, exist := existMap[nums2[i]]; exist {
			res = append(res, nums2[i])
			delete(existMap, nums2[i])
		}
	}

	return res
}

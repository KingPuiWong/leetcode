package stackqueue

import "sort"

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}

	var ans []int
	for k := range countMap {
		ans = append(ans, k)
	}

	sort.Slice(ans, func(i, j int) bool {
		return countMap[ans[i]] > countMap[ans[j]]
	})

	return ans[:k]
}

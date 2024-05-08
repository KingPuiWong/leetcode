package string

/*
给你两个字符串 haystack 和 needle ，

请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。

如果 needle 不是 haystack 的一部分，则返回  -1 。
*/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		tempI := i
		for j := 0; j < len(needle); j++ {
			if tempI >= len(haystack) {
				return -1
			}
			if haystack[tempI] == needle[j] {
				if j == len(needle)-1 {
					return i
				}
				tempI++
			} else {
				break
			}
		}
	}
	return -1
}

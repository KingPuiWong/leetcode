package hash

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
*/
func canConstruct(ransomNote string, magazine string) bool {
	magMap := make(map[int32]int)
	for _, v := range magazine {
		magMap[v-'a']++
	}

	for _, v := range ransomNote {
		magMap[v-'a']--
		if magMap[v-'a'] < 0 {
			return false
		}
	}
	return true
}

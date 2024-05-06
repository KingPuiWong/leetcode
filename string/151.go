package string

/*
*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/
func reverseWords(s string) string {
	// 删除前面、后面和中间多余的空格
	b := []byte(s)
	var slow int
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}

			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}

	}

	b = b[0:slow]

	// 翻转整个字符串
	revert(b)
	// 翻转单个单词
	var last int
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {

			revert(b[last:i])
			last = i + 1
		}
	}

	return string(b)

}

func revert(s []byte) {
	startIndex := 0
	endIndex := len(s) - 1
	for startIndex < endIndex {
		s[startIndex], s[endIndex] = s[endIndex], s[startIndex]
		startIndex++
		endIndex--
	}
}

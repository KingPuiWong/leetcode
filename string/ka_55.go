package string

import (
	"fmt"
)

/**
题目描述:
字符串的右旋转操作是把字符串尾部的若干个字符转移到字符串的前面。
给定一个字符串 s 和一个正整数 k，
请编写一个函数，将字符串中的后面 k 个字符移到字符串的前面，
实现字符串的右旋转操作。
例如，对于输入字符串 "abcdefg" 和整数2，
函数应该将其转换为 "fgabcde"。

输入描述
输入共包含两行，第一行为一个正整数 k，
代表右旋转的位数。第二行为字符串 s，代表需要旋转的字符串。

输出描述
输出共一行，为进行了右旋转操作后的字符串。
*/

func ka() {
	var str string
	var target int

	fmt.Scanln(&target)
	fmt.Scanln(&str)

	var right []byte
	var left []byte
	for i := 0; i < len(str); i++ {
		if i < len(str)-target {
			left = append(left, str[i])
		} else {
			right = append(right, str[i])
		}
	}

	var res []byte
	res = append(res, append(right, left...)...)
	fmt.Printf(string(res))

}

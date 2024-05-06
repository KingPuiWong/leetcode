package string

import (
	"fmt"
)

/**
给定一个字符串 s，它包含小写字母和数字字符，
请编写一个函数，将字符串中的字母字符保持不变，
而将每个数字字符替换为number。
例如，对于输入字符串 "a1b2c3"，
函数应该将其转换为 "anumberbnumbercnumber"。
*/

func main() {
	var strByte []byte
	fmt.Scanln(&strByte)
	for i := 0; i < len(strByte); i++ {
		if strByte[i] <= '9' && strByte[i] >= '0' {
			insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strByte = append(strByte[:i], append(insertElement, strByte[i+1:]...)...)
			i = i + len(insertElement) - 1
		}
	}

	fmt.Printf(string(strByte))
}

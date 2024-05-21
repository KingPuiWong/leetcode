package string

import "strings"

func repeatedSubstringPattern(s string) bool {
	joinString := s + s
	return strings.Contains(joinString[1:len(joinString)-1], s)
}

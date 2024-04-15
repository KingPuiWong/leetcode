package hash

func isAnagram(s string, t string) bool {
	firstHashMap := make(map[string]int)
	secondHashMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		firstHashMap[string(s[i])]++
	}

	for i := 0; i < len(t); i++ {
		secondHashMap[string(t[i])]++
	}
	if len(s) > len(t) {
		for key, value := range firstHashMap {
			if secondHashMap[key] != value {
				return false
			}
		}
	} else {
		for key, value := range secondHashMap {
			if firstHashMap[key] != value {
				return false
			}
		}
	}
	return true
}

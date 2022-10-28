package ch2_q2_first_unique_char

// firstUniqChar ...
func firstUniqChar(s string) int {
	uniqueMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := uniqueMap[s[i]]; ok {
			uniqueMap[s[i]] += 1
		} else {
			uniqueMap[s[i]] = 1
		}
	}
	for j := 0; j < len(s); j++ {
		if uniqueMap[s[j]] == 1 {
			return j
		}
	}
	return -1
}

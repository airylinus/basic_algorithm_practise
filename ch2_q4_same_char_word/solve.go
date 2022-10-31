package ch2_q4_same_char_word

func isSameCharWord(source, target string) bool {
	if len(source) != len(target) {
		return false
	}
	charMap := map[byte]int{}
	for i := 0; i < len(source); i++ {
		if _, ok := charMap[source[i]]; ok {
			charMap[source[i]] += 1
		} else {
			charMap[source[i]] += 1
		}
		if _, ok := charMap[target[i]]; ok {
			charMap[target[i]] -= 1
		} else {
			charMap[target[i]] -= 1
		}
	}
	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}
	return true
}

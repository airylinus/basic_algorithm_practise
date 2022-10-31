package ch2_q9_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	r := []byte{}
	for n := 0; n < len(strs[0]); n++ {
		currentByte := strs[0][n]
		for k, _ := range strs {
			if len(strs[k]) <= n {
				return string(r)
			}
			if strs[k][n] != currentByte {
				return string(r)
			}
		}
		r = append(r, currentByte)
	}
	return string(r)
}

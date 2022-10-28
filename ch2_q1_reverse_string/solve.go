package ch2_q1_reverse_string

func reverseString(s []byte) (r []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
	return s
}

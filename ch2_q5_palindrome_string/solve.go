package ch2_q5_palindrome_string

import "strings"

/*
 * 0 ==> 48
 * 9 ==> 57
 * 65 ==> A
 * 90 ==> Z
 * 97 ==> a
 * 122 ==> z
 */

func palindromeString(target string) bool {
	target = strings.ToLower(target)
	h, e := 0, len(target)-1
	for h <= e {
		if !isByteCharNumber(target[h]) {
			h += 1
			continue
		}
		if !isByteCharNumber(target[e]) {
			e -= 1
			continue
		}
		if target[h] != target[e] {
			return false
		}
		h += 1
		e -= 1
	}
	return true
}

func isByteCharNumber(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	if b >= 65 && b <= 90 {
		return true
	}
	if b >= 97 && b <= 122 {
		return true
	}
	return false
}

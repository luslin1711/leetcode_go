package _59_repeatedSubstringPattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}
	t := s + s
	return strings.Index(t[1:len(t)-1],s) != -1
}
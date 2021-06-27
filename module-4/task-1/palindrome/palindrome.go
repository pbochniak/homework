package palindrome

import "unicode"

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}
	}
	return true
}

package lengthOfLongestSubstring

func hasRepeatingChars(str string) bool {
	chars := make(map[rune]bool)
	for _, s := range str {
		if chars[s] {
			return true
		}
		chars[s] = true
	}
	return false
}

func LengthOfLongestSubstring(s string) int {
	longestSubstingLen := 0
	substrings := make(map[string]bool)

	for leftPtr, rightPtr := 0, 0; rightPtr < len(s); rightPtr++ {
		substr := s[leftPtr : rightPtr+1]
		if !substrings[substr] && hasRepeatingChars(substr) {
			leftPtr += 1
		} else if !substrings[substr] {
			substrings[substr] = true
			substrLen := len(substr)
			subsLen := substrLen
			if longestSubstingLen < substrLen {
				longestSubstingLen = subsLen
			}
		}
	}

	return longestSubstingLen
}

package longestPalindrome

func LongestPalindrome(s string) string {
	sLen := len(s)
	lpLen := 0

	lp := ""
	if sLen > 0 {
		lp = s[0:1]
		lpLen = 1
	}

	for leftPtr := 0; sLen-leftPtr > lpLen; leftPtr++ {
		isPalindrome := true
		for rightPtr := sLen; rightPtr-leftPtr > lpLen; rightPtr-- {
			substring := s[leftPtr:rightPtr]

			ssLen := len(substring)
			halfSubstring := ssLen / 2

			for iL := 0; iL < halfSubstring; iL++ {
				isPalindrome = substring[iL] == substring[ssLen-iL-1]
				if !isPalindrome {
					break
				}
			}

			if isPalindrome {
				if len(lp) < len(substring) {
					lp = substring
					lpLen = len(lp)
				}
				break
			}
		}
	}

	return lp
}

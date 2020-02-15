package palindromes

func isPalindrome(s string) bool {
	sLen := len(s)
	if sLen > 0 {
		for i := 0; i <= sLen/2; i++ {
			if s[i] != s[sLen-i-1] {
				return false
			}
		}
	}
	return true
}

func ShortestPalindrome(s string) string {
	sp := s
	ind := len(s)
	for ind > 1 && !isPalindrome(s[0:ind]) {
		ind--
	}

	for !isPalindrome(sp) {
		sp = s[ind:ind+1] + sp
		ind++
	}

	return sp
}

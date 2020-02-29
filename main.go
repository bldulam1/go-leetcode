package main

import "fmt"

func main() {
	inputs := map[string]string{
		"1":   "0",
		"11":  "9",
		"88":  "77",
		"230": "232",
		"100": "99",
		"200": "202",
		"40":  "44",
	}

	for test, expect := range inputs {
		np := nearestPalindromic(test)
		if np != expect {
			fmt.Println("FAILED", test, np)
		} else {
			fmt.Println("passed", test, np)
		}
	}
}

func isPalindrome(s []rune) bool {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}

func nearestPalindromic(n string) string {
	nrune := []rune(n)
	if n == "0" {
		return n
	}

	if isPalindrome(nrune) {
		isBorrowing := true
		for i := len(n) / 2; i >= 0; i-- {
			if isBorrowing {
				nrune[i]--
			} else {
				break
			}

			if nrune[i] < '0' {
				nrune[i] = '9'
			} else {
				isBorrowing = false
			}
		}
	}

	for nrune[0] == '0' && len(nrune) > 1 {
		nrune = nrune[1:]
	}

	for i := 0; i < len(nrune)/2; i++ {
		nrune[len(nrune)-1-i] = nrune[i]
	}

	return string(nrune)
}

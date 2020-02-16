package validNumber

func IsNumber(s string) bool {
	sCopy := s
	if len(sCopy) > 1 {
		for sCopy[0] == ' ' {
			sCopy = sCopy[1:]
		}
		for sCopy[len(sCopy)-1] == ' ' {
			sCopy = sCopy[:len(sCopy)-1]
		}
	}

	if sCopy[0] < '0' || sCopy[0] > '9' {
		return false
	}

	isEexists := false
	isDotExists := false
	for ind, char := range sCopy {
		if '0' <= char && char <= '9' {
			continue
		} else if char == '+' || char == '-' {
			if ind == 0 {
				if ind < len(sCopy)-1 && (sCopy[ind+1] == '.' || sCopy[ind+1] == 'e') {
					return false
				}
			} else if sCopy[ind-1] != 'e' && !(sCopy[ind-1] >= '0' && sCopy[ind-1] <= '9') ||
				sCopy[ind-1] >= '0' && sCopy[ind-1] <= '9' ||
				ind == len(sCopy)-1 {
				return false
			}
		} else if char == 'e' {
			if isEexists || ind == 0 || ind == len(sCopy)-1 || sCopy[ind-1] == '.' {
				return false
			} else {
				isEexists = true
			}
		} else if char == '.' {
			if isEexists || isDotExists {
				return false
			}
			isDotExists = true
		} else if char == ' ' {
			return false
		} else {
			return false
		}
	}

	return true
}

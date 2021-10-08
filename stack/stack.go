package stack

func IsValidClosure(input string) bool {
	s := make([]rune, 0)
	pairMap := make(map[rune]rune, 0)
	pairMap['{'] = '}'
	pairMap['['] = ']'
	pairMap['('] = ')'

	for _, chr := range input {
		if _, ok := pairMap[chr]; ok {
			s = append(s, chr)
		} else if len(s) > 0 {
			opener := s[len(s)-1]
			s = s[:len(s)-1]
			if pairMap[opener] != chr {
				return false
			}
		} else {
			return false
		}
	}

	return len(s) == 0
}

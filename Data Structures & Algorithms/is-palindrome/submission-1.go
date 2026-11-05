func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !((s[l] >= 'a' && s[l] <= 'z') || (s[l] >= 'A' && s[l] <= 'Z') || (s[l] >= '0' && s[l] <= '9')) {
			l++
		}
		for l < r && !((s[r] >= 'a' && s[r] <= 'z') || (s[r] >= 'A' && s[r] <= 'Z') || (s[r] >= '0' && s[r] <= '9')) {
			r--
		}

		cl, cr := s[l], s[r]
		if cl >= 'A' && cl <= 'Z' {
			cl += 'a' - 'A'
		}
		if cr >= 'A' && cr <= 'Z' {
			cr += 'a' - 'A'
		}

		if cl != cr {
			return false
		}

		l++
		r--
	}

	return true
}
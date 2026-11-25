func lengthOfLongestSubstring(s string) int {
	n := len(s)
	l := 0
	sw := make(map[byte]int)
	maxx := 0

	for r := 0; r < n; r++ {
		if _, ok := sw[s[r]]; ok {
			if sw[s[r]] >= l {
				l = sw[s[r]] + 1
			}
		}
		sw[s[r]] = r

		maxx = max(maxx, r-l+1)
	}

	return maxx
}
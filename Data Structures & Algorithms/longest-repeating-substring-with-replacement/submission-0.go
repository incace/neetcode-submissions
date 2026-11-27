func characterReplacement(s string, k int) int {
	var count [26]int
	left := 0
	res := 0

	maxfreq := func() int {
		maxF := 0
		for _, c := range count {
			if c > maxF {
				maxF = c
			}
		}
		return maxF
	}

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++

		for (right-left+1)-maxfreq() > k {
			count[s[left]-'A']--
			left++
		}

		if curLen := right - left + 1; curLen > res {
			res = curLen
		}
	}

	return res
}
func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := piles[0]
	for _, p := range piles[1:] {
		if p > r {
			r = p
		}
	}

	res := r
	for l <= r {
		k := l + (r-l)/2

		hours := 0
		for _, p := range piles {
			hours += (p + k - 1) / k
		}

		if hours <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}
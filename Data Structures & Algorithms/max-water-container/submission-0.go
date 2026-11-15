func maxArea(heights []int) int {
	l, r := 0, len(heights) -1
	maxArea := 0

	for l < r {
		curl, curr := heights[l], heights[r]
		maxArea = max(maxArea, (r - l) * min(curl, curr))
		if curl > curr {
			for curr >= heights[r] && l < r{
				r--
			}
		} else {
			for curl >= heights[l] && l < r{
				l++
			}
		}
	}


	return maxArea
}

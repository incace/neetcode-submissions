type Pair struct {
    Index int
	Height int
}


func largestRectangleArea(heights []int) int {
	stack := make([]Pair, len(heights))
	sp := -1
	maxArea := 0

	for i, h := range heights{
		backEx := i
		if sp >= 0{
			for sp >= 0 && stack[sp].Height >= h{
				maxArea = max(maxArea, stack[sp].Height * (i - stack[sp].Index))
				backEx = stack[sp].Index
				sp--
			}
		}
		sp++
		stack[sp] = Pair{backEx, h}
	}

	for sp >= 0{
		maxArea = max(maxArea, stack[sp].Height * (len(heights) - stack[sp].Index))
		sp--
	}

	return maxArea
}

type drop struct {
	i int
	h int
}

func trap(height []int) int {
	res := 0
	n := len(height)
	stack := make([]drop, n)
	sp := 0

	for cur := 0; cur < n; cur++ {
		for sp > 0 && height[cur] > stack[sp-1].h {
			top := &stack[sp-1]
			if height[top.i] <= height[cur] {
				res += (height[top.i] - top.h) * (cur - top.i - 1)
				sp--
			} else {
				res += (height[cur] - top.h) * (cur - top.i - 1)
				top.h = height[cur]
				break
			}
		}
		if cur+1 < n && height[cur] > height[cur+1] {
			stack[sp] = drop{i: cur, h: height[cur+1]}
			sp++
		}
	}
	return res
}

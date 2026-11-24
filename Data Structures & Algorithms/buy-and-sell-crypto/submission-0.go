func maxProfit(prices []int) int {
	res := 0
	cm := 0

	for r:= 0; r < len(prices); r++{
		if prices[cm] > prices[r]{
			cm = r
		}
		res = max(res, prices[r] - prices[cm])
	}
	return res
}

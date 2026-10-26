func evalRPN(tokens []string) int {
	var stack []int
	n := 0

	var res int

	for _, r := range tokens {
		switch r {
		case "+":
			res = stack[n-1] + stack[n-2]
			stack = stack[:n-2]
			stack = append(stack, res)
			n--
		case "-":
			res = stack[n-2] - stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, res)
			n--
		case "*":
			res = stack[n-2] * stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, res)
			n--
		case "/":
			res = stack[n-2] / stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, res)
			n--
		default:
			res, _ = strconv.Atoi(r)
			stack = append(stack, res)
			n++
		}
	}
	return res
}

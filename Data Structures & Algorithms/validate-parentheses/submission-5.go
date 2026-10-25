func isValid(s string) bool {
    stack := make([]byte, 0)
	n := 0

	if len(s) % 2 == 1 {
		return false
	}
	for i, _ := range s{
		ch := s[i]
		if ch == '[' || ch == '(' || ch == '{'{
			stack = append(stack, ch)
			n++
		} else {
			if len(stack) == 0{
				return false
			}
			pop_val := stack[n-1]
			stack = stack[:n-1]
			n--

			if ch == ']' && pop_val != '['{
				return false
			} 
			if ch == ')' && pop_val != '('{
				return false
			} 
			if ch == '}' && pop_val != '{'{
				return false
			}
		}
	}

	if len(stack) > 0{
		return false
	}
	return true
}

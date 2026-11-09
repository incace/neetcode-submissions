func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n - 1
	tmp := numbers[l] + numbers[r]
	for tmp != target{
		if tmp > target {
			r--;
		} else{
			l++;
		}
		tmp = numbers[l] + numbers[r]
	}

	return []int{l + 1, r + 1}
}

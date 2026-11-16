func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		cur_ind := (l + r) / 2
		if nums[cur_ind] == target{
			return cur_ind
		}

		if nums[cur_ind] > target{
			r = cur_ind - 1
		} else{
			l = cur_ind + 1
		}
	}

	return -1
}
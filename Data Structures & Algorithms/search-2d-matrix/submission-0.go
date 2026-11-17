func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	l, r := 0, n*m - 1

	for l <= r {
		cur_ind := l + (r - l) / 2
		cur_m := cur_ind / n
		cur_n := cur_ind % n
		if matrix[cur_m][cur_n] == target{
			return true
		}

		if matrix[cur_m][cur_n] > target{
			r = cur_ind - 1
		} else{
			l = cur_ind + 1
		}
	}

	return false
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
    cols := make([]map[byte]struct{}, 9)
    squares := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]struct{})
        cols[i] = make(map[byte]struct{})
        squares[i] = make(map[byte]struct{})
    }
	
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++{

			if board[i][j] == '.' { continue }
			if _, ok := rows[i][board[i][j]]; ok{
				return false
			}
			rows[i][board[i][j]] = struct{}{}

			if _, ok := cols[j][board[i][j]]; ok{
				return false
			}
			cols[j][board[i][j]] = struct{}{}


			sq_num := (i / 3) * 3 + (j / 3)
			if _, ok := squares[sq_num][board[i][j]]; ok{
				return false
			}
			squares[sq_num][board[i][j]] = struct{}{}

		}
	}
	return true
}

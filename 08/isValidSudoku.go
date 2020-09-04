package main
//43 有效的数独 判断数独是否有效
func isValidSudoku(board [][]byte) bool {
	var col, row, sBox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				boxIndex := (i/3)*3+j/3
				//列中是否存在该数字
				if col[i][num] == 1 {
					return false
				} else {
					col[i][num] = 1
				}
				//行中是否存在该数字
				if row[j][num] == 1 {
					return false
				} else {
					row[j][num] = 1
				}
				//本块是否存在该数字
				if sBox[boxIndex][num] == 1 {
					return false
				} else {
					sBox[boxIndex][num] = 1
				}
			}
		}
	}
	return true
}
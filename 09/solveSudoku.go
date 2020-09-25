package main

//37. 解数独

//backtrace
func solveSudoku(board [][]byte) {
	//去除非法数独
	if board == nil || len(board) != 9 {
		return
	}
	for i := 0; i < 9; i++ {
		if board[i] == nil || len(board[i]) != 9 {
			return
		}
	}
	//回溯
	backTrack(board, 0, 0)
}

func backTrack(board [][]byte, row int, col int) bool {
	if col == 9 {
		return backTrack(board, row+1, 0)
	}
	if row == 9 {
		return true
	}

	if board[row][col] != '.' {
		return backTrack(board, row, col+1)
	}

	for num := byte('1'); num <= '9'; num++ {
		if !isValidNumber(board, row, col, num) {
			continue
		}
		board[row][col] = num
		if backTrack(board, row, col+1) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func isValidNumber(board [][]byte, row int, col int, number byte) bool {
	// 三个方向，任一方向重复，ch就不能放在这个位置
	for k := 0; k < 9; k++ {
		// 同一行九个位置已出现 ch
		if board[row][k] == number {
			return false
		}
		// 同一列九个位置中已出现 ch
		if board[k][col] == number {
			return false
		}
		// 同一个子数独九个位置中已出现 ch
		if board[(row/3)*3+k/3][(col/3)*3+k%3] == number {
			return false
		}
	}
	return true
}

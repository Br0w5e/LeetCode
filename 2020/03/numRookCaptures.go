// 计算车可捕获的数量
func numRookCaptures(board [][]byte) int {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
			return cap(board, i, j, 0, 1) + cap(board, i, j, 0, -1) + cap(board, i, j, 1, 0) + cap(board, i, j, -1, 0)
			}
		}
	}
	return 0
}

func cap(board [][]byte, x, y, dx, dy int) int {
	for x >= 0 && x < 8 && y >= 0 && y < 8 && board[x][y] != 'B' {
		if board[x][y] == 'p' {
			return 1
		}
		x += dx
		y += dy
	}
	return 0
}
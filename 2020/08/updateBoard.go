package main
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}
//点击一次
func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	//点击直接爆炸
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(board, x, y)
	}
	return board
}

func dfs(board [][]byte, x, y int)  {
	cnt := 0
	//试探点击点周围雷的个数
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}
	//周围有雷，直接返回周围雷的个数
	if cnt > 0 {
		board[x][y] = byte(cnt+'0')
	} else {
		//周围没有雷，该位置标B并进行周围判断，直到有数字
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			dfs(board, tx, ty)
		}
	}
}

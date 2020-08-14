package main

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)
//bfs
func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	n, m := len(board), len(board[0])
	queue := [][]int{}
	for i := 0; i < n; i++ {
		//左边
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		//右边
		if board[i][m-1] == 'O' {
			queue = append(queue, []int{i, m-1})
		}
	}
	for i := 1; i < m-1; i++ {
		//上边
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
		}
		//下边
		if board[n-1][i] == 'O' {
			queue = append(queue, []int{n-1, i})
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx > 0 && mx < n && my > 0 && my < m && board[mx][my] == 'O' {
				queue = append(queue, []int{mx, my})
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}


//dfs
func solve1(board [][]byte)  {
	x := len(board)
	if x == 0 {
		return
	}
	y := len(board[0])
	for i := 0; i < x; i ++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][y-1] == 'O' {
			dfs(board, i, y-1)
		}
	}
	for i := 1; i < y; i ++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[x-1][i] == 'O' {
			dfs(board, x-1, i)
		}
	}

	for i := 0; i < x; i ++ {
		for j := 0; j < y; j ++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >=len(board[0]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '#'
		dfs(board, i-1, j)
		dfs(board, i+1, j)
		dfs(board, i, j+1)
		dfs(board, i, j-1)
	}
}
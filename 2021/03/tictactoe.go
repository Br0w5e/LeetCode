package main

//面试题 16.04. 井字游戏

func tictactoe(board []string) string {
	N := len(board)
	row := make([]int, N)
	col := make([]int, N)
	dia1 := 0 //主对角线
	dia2 := 0 //副对角线
	pending := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			offset := 0
			if board[i][j] == 'O' {
				offset = 1
			} else if board[i][j] == 'X' {
				offset = -1
			} else {
				pending = 1
			}
			row[i] += offset
			col[j] += offset

			if i == j {
				dia1 += offset
			}
			if i+j == N-1 {
				dia2 += offset
			}
		}
	}
	if dia1 == N || dia2 == N {
		return "O"
	}
	if dia1 == -N || dia2 == -N {
		return "X"
	}
	for k := 0; k < N; k++ {
		if row[k] == N || col[k] == N {
			return "O"
		} else if row[k] == -N || col[k] == -N {
			return "X"
		}
	}
	if pending == 0 {
		return "Draw"
	}
	return "Pending"
}

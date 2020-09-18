package main

//79 单词搜索

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, x int, y int, index int) bool {
	var dx = [4]int{-1, 1, 0, 0}
	var dy = [4]int{0, 0, -1, 1}

	if index == len(word) {
		return true
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[index] {
		return false
	}
	tmp := board[x][y]
	board[x][y] = 0
	for k := 0; k < 4; k++ {
		if search(board, word, x+dx[k], y+dy[k], index+1) {
			return true
		}
	}
	// if search(board, word, i+1, j, k+1) || search(board, word, i-1, j, k+1) || search(board, word, i, j+1, k+1) || search(board, word, i, j-1, k+1){
	//     return true
	// }
	board[x][y] = tmp
	return false
}
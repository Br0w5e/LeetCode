package main

//1314. 矩阵区域和

func matrixBlockSum(mat [][]int, K int) [][]int {
	sum := getInitValue(K, mat)
	r := make([][]int, len(mat))
	for i := 0; i < len(r); i++ {
		r[i] = make([]int, len(mat[i]))
		r[i][0] = sum
		for j := 1; j < len(mat[i]); j++ {
			r[i][j] = moveX(r[i][j-1], mat, j, i, K)
		}
		sum = moveY(r[i][0], mat, 0, i+1, K)
	}
	return r
}

func getInitValue(K int, mat [][]int) int {
	sum := 0
	for i := 0; i <= K; i++ {
		if i >= len(mat) {
			continue
		}
		for j := 0; j <= K; j++ {
			if j >= len(mat[i]) {
				continue
			}
			sum += mat[i][j]
		}
	}
	return sum
}

func moveX(sum int, mat [][]int, x, y int, K int) int {
	if x + K < len(mat[0]) {
		for i := y - K; i <= y + K; i++ {
			if i < 0 || i >= len(mat) {
				continue
			}
			sum += mat[i][x+K]
		}
	}
	if x - K - 1 >= 0 {
		for i := y - K; i <= y + K; i++ {
			if i < 0 || i >= len(mat) {
				continue
			}
			sum -= mat[i][x-K-1]
		}
	}
	return sum
}

func moveY(sum int, mat [][]int, x, y int, K int) int {
	if y + K < len(mat) {
		for i := x - K; i <= x + K; i++ {
			if i < 0 || i >= len(mat[0]) {
				continue
			}
			sum += mat[y+K][i]
		}
	}
	if y - K - 1 >= 0 {
		for i := x - K; i <= x + K; i++ {
			if i < 0 || i >= len(mat[0]) {
				continue
			}
			sum -= mat[y-K-1][i]
		}
	}
	return sum
}

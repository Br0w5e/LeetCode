package main

func imageSmoother(M [][]int) [][]int {
	R, C := len(M), len(M[0])
	res := make([][]int, R)
	for i := range res {
		res[i] = make([]int, C)
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			cnt := 0
			for r := i-1; r <= i+1; r++ {
				for c := j-1; c <= j+1; c++ {
					if 0 <= r && r < R && 0 <= c && c < C {
						res[i][j] += M[r][c]
						cnt++
					}
				}
			}
			res[i][j] /= cnt
		}
	}
	return res
}

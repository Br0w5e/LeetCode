package main

//52. N皇后 II

//不太懂

func totalNQueens(n int) int {
	count := 0
	return dfs(n, count, 0, 0, 0, 0)
}

func dfs(n, count, row, col, pie, na int) int {
	if row >= n {
		count++
		return count
	}

	bits := (^(col|pie|na)) & ((1<<n)-1)
	for bits > 0 {
		p := bits & -bits
		count = dfs(n, count, row+1, col|p, (pie|p)<<1, (na|p)>>1)
		bits = bits & (bits-1)
	}
	return count
}
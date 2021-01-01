package main

//861. 翻转矩阵后的得分

//首列全部置为1（行变化）， 后边的统计每列1的数目大于0的数目（列变换）
func matrixScore(A [][]int) int {
	n, m := len(A), len(A[0])
	res := 1 <<(m-1)*n
	for i := 1; i < m; i++ {
		one, zero := 0, 0
		for _, row := range A {
			if row[i] == row[0] {
				one++
			} else {
				zero++
			}
		}
		if one < zero {
			one = zero
		}
		res += 1 << (m-1 - i)*one
	}
	return res
}

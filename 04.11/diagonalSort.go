//将对角线上面的元素进行排列
//弄清对角线元素关系就行[i,j] [i+k,j+k]
func diagonalSort(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			for k := 0; i+k < n && j+k < m; k++{
				if mat[i][j] > mat[i+k][j+k]{
					mat[i][j], mat[i+k][j+k] = mat[i+k][j+k], mat[i][j]
				}
			}
		}
	}
	return mat
}
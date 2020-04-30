//将矩阵进行顺时针选择90度
//重点： 推导旋转前后的坐标关系
//matrix[i][j]=matrix[j][n-1-i]


//方法1：先翻转，再对折
//翻转：matrix[i][j] = matrix[n-1-i][j]
//对折：matrix[i][j] = matrix[j][i]

func rotate(matrix [][]int)  {
	n := len(matrix)
	//翻转
	for i := 0; i < n/2; i++{
		for j := 0; j < n; j++{
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	//对折
	for i := 0; i < n-1; i++{
		for j := i+1; j < n; j++{
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return 
}

//原地旋转：每次动四个点，保证不覆盖
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j <= n-i-2; j++{
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]	
		}
	}
	return
}
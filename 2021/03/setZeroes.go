package main

//73. 矩阵置零
//面试题 01.08. 零矩阵

//将矩阵0所在的行和列都置位0
//方法：求出0所在行和列的行列号，然后进行求解
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	grid := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				tmp := make([]int, 2)
				tmp[0], tmp[1] = i, j
				grid = append(grid, tmp)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < len(grid); k++ {
				if i == grid[k][0] || j == grid[k][1] {
					matrix[i][j] = 0
				}
			}
		}
	}
	return
}

//先标记，然后处理
func setZeroes(matrix [][]int) {
	tags := make([][]int, 0)
	m, n := len(matrix), len(matrix[0])
	//寻找为0位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				tags = append(tags, []int{i, j})
			}
		}
	}
	//进行修改
	for _, tag := range tags {
		for i := 0; i < m; i++ {
			matrix[i][tag[1]] = 0
		}
		for j := 0; j < n; j++ {
			matrix[tag[0]][j] = 0
		}
	}
}

package main

//1582. 二进制矩阵中的特殊位置

//稀疏矩阵类似方法
func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	nums := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				nums = append(nums, []int{i, j})
			}
		}
	}
	res := 0
	for _, num := range nums {
		//标记
		flag := false
		//同一列
		for i := 0; i < m; i++ {
			if mat[i][num[1]] == 1 && i != num[0] {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		//同一行
		for i := 0; i < n; i++ {
			if mat[num[0]][i] == 1 && i != num[1] {
				flag = true
				break
			}
		}
		if !flag {
			res++
		}
	}
	return res
}

//记录行之和和列之和
func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	//列和
	cols := make([]int, n)
	//行和
	rows := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rows[i] += mat[i][j]
			cols[j] += mat[i][j]
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && cols[j] == 1 && rows[i] == 1 {
				res++
			}
		}
	}
	return res
}
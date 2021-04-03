package main

//74. 搜索二维矩阵
//面试题 10.09. 排序矩阵查找

//方法1：暴力搜索
func searchMatrix1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, num := range row {
			if num == target {
				return true
			}
		}
	}
	return false
}

//方法2：二维降一维进行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		r, c := mid/n, mid%n
		if matrix[r][c] < target {
			left = mid + 1
		} else if matrix[r][c] == target {
			return true
		} else {
			right = mid - 1
		}
	}
	return false
}

//对角线遍历（右上到左下的对角线）
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1
	for r < m && c >= 0 {
		if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		} else {
			return true
		}
	}
	return false
}

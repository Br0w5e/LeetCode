package main

//74. 搜索二维矩阵

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
	if len(matrix) == 0 ||  len(matrix[0]) == 0{
		return false
	}
	n, m := len(matrix), len(matrix[0])

	if target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}
	left, right := 0, n*m-1
	for left <= right {
		mid := left + (right-left)>>1
		i := mid / m
		j := mid % m
		if matrix[i][j] == target {
			return true
			//中值大于目标，降低右边
		} else if matrix[i][j] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return false
}
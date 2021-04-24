package main

//363. 矩形区域不超过 K 的最大数值和
import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	//计算列的前缀和
	rowNum, colNum := len(matrix), len(matrix[0])
	for row := 0; row < rowNum; row++ {
		for col := 1; col < colNum; col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}

	res := math.MinInt32
	// 固定左右边界，移动上下边界，求和
	for left := 0; left < colNum; left++ {
		for right := left; right < colNum; right++ {
			for top := 0; top < rowNum; top++ {
				sum := 0
				for bottom := top; bottom < rowNum; bottom++ {
					if left == 0 {
						//左边界为0则不用减
						sum += matrix[bottom][right]
					} else {
						//右边界前缀-左边界前缀，并且自上而下求和。得到左上角为[top,left]，右下角为[bottom, right]的矩阵面积
						sum += matrix[bottom][right] - matrix[bottom][left-1]
					}
					//每次都得比较的
					if sum <= k && sum > res {
						res = sum
					}
				}
			}
		}
	}
	return res
}

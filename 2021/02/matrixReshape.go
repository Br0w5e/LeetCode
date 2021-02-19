package main
//566. 重塑矩阵

//重塑矩阵
func matrixReshape(nums [][]int, r int, c int) [][]int {
	rw, cw := len(nums), len(nums[0])
	if rw*cw != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	k := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			//都是行宽
			res[i][j] = nums[k/cw][k%cw]
			k++
		}
	}
	return res
}
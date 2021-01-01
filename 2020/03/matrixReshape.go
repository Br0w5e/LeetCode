//重塑矩阵
func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	m := len(nums[0])
	if n*m != r*c{
		return nums
	}
	tmp := make([]int, 0)
	for i := 0; i < n; i++{
		for j := 0; j < m; j++ {
			tmp = append(tmp, nums[i][j])
		}
	}
	res := make([][]int, 0)
	for i := 0; i < r; i++{
		res = append(res, nums[i*c:(i+1)*c])
	}
	return res
}
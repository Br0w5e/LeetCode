// 统计变换后的数组中奇数的个数
func oddCells(n int, m int, indices [][]int) int {
	zero := make([]int, m)
	nums := make([][]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, zero)
	}
	for i := 0; i < len(indices); i++ {
		x := indices[i][0]
		for j := 0; j < m; j++ {
			nums[x][j]++
		}
		y := indices[i][1]
		for k := 0; k < n; k++ {
			nums[k][y]++
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums[i][j]%2 == 1 {
				res++
			}
		}
	}
	return res
}

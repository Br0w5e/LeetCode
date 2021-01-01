package main

//1260. 二维网格迁移
//暴力模拟

func shiftGrid(grid [][]int, k int) [][]int {
	tmp := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			tmp = append(tmp, grid[i][j])
		}
	}
	k %= len(tmp)
	for w := 0; w < k; w++ {
		tmp_V := 0
		for i := len(tmp)-1; i >= 0; i-- {
			if i == len(tmp)-1 {
				tmp_V = tmp[i]
				tmp[i] = tmp[i-1]
			} else if i > 0{
				tmp[i] = tmp[i-1]
			} else {
				tmp[i] = tmp_V
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = tmp[i*len(grid[0])+j]
		}
	}
	return grid
}

//三次翻转
func shiftGrid1(grid [][]int, k int) [][]int {
	nums := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		nums = append(nums, grid[i]...)
	}
	k %= len(nums)
	n := len(nums)
	//前n-k翻转
	reverse(nums[:n-k])
	//后n-k翻转
	reverse(nums[n-k:])
	//集体翻转
	reverse(nums)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = nums[i*len(grid[0])+j]
		}
	}
	return grid
}

func reverse(nums []int) {  // 翻转数组
	left, right := 0, len(nums)-1
	for left < right {
		nums[left],nums[right] = nums[right],nums[left]
		left++
		right--
	}
}

//直接操作
func shiftGrid2(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	if k == n*m {
		return grid
	}
	nums := make([]int, 0)
	for _, row := range grid {
		nums = append(nums, row...)
	}
	k %= len(nums)

	for i, v := range nums {
		//体现旋转
		col, row := ((i+k)/m)%n, (i+k)%m
		grid[col][row] = v
	}
	return grid
}
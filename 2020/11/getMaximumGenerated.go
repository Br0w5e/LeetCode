package main
//5561. 获取生成数组中的最大值
import "sort"

//模拟
//没读懂题

func getMaximumGenerated(n int) int {
	//corner case
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[0], nums[1] = 0, 1
	for i := 0; i < n-1; i++ {
		if 2*i >= 2 && 2*i <= n {
			nums[2*i] = nums[i]
		}
		if 2*i+1 >= 2 && 2*i+1 <= n {
			nums[2 * i + 1] = nums[i] + nums[i + 1]
		}
	}
	sort.Ints(nums)
	return nums[n]
}

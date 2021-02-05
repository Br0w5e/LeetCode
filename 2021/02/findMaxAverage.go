package main
//643. 子数组最大平均数 I

//暴力
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0.0
	}
	var res float64 = -10000.0
	for i := 0; i < len(nums)-k+1; i++ {
		res = max(Sum(nums[i:i+k]), res)
	}
	return res
}

func Sum(nums []int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum)/float64(len(nums))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}


//滑动窗口
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, num := range nums[:k] {
		sum += num
	}
	maxSum := sum
	for i := 0; i < len(nums)-k; i++ {
		sum += (nums[i+k]-nums[i])
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum)/float64(k)
}



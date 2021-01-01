package main

//5610. 有序数组中差绝对值之和
//注意推关系式
func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[0] += (nums[i]-nums[0])
	}

	for i := 1; i < n; i++ {
		//res[i] = res[i-1] + (nums[i]-nums[i-1])*(i-1) + (nums[i-1]-nums[i]) * (n-1 - i)
		res[i] = res[i-1] + (nums[i]-nums[i-1])*(2*i-n)
	}
	return res
}

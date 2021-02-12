package main

//5672. 检查数组是否经排序和轮转得到

//模拟
func check(nums []int) bool {
	flag := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			flag++
		}
		//标记点大于等于两个，返回false
		if flag >= 2 {
			return false
		}
	}
	//标记点等于0返回true，或者标记点等于1且最后的小于等于第一个
	if flag == 0 {
		return true
	} else {
		return nums[len(nums)-1] <= nums[0]
	}
}


//好一点的做法
func check(nums []int) bool {
	flag := 0
	for i := 0; i < len(nums); i++ {
		//点睛之笔
		if nums[i] > nums[(i+1)%len(nums)] {
			flag++
		}
		//标记点大于等于两个，返回false
		if flag > 1 {
			return false
		}
	}
	return true
}
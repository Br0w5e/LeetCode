package main

//旋转数组
//将数组顺时针旋转k次
//朴素方法：模拟
func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	if k == 0 {
		return nums
	}
	for i := 0; i < k; i++ {
		tmp := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
	return nums
}

//三次翻转：
/*
原始数组                  : 1 2 3 4 5 6 7
反转所有数字后             : 7 6 5 4 3 2 1
反转前 k 个数字后          : 5 6 7 4 3 2 1
反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
*/

func reverse(nums []int) { // 翻转数组
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums)
	//前n-k翻转
	reverse(nums[:k])
	//后n-k翻转
	reverse(nums[k:])
	//集体翻转
}

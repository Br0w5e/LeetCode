package main

//5719. 每个查询的最大异或值

import "fmt"

//暴力--->超时
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := n; i > 0; i-- {
		res[n-i] = help(nums[:i], maximumBit)
	}
	return res
}

func help(nums []int, maximumBit int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	tmp := make([]int, 0)
	k := 1
	for k < 1<<maximumBit {
		tmp = append(tmp, 1-xor&1)
		xor >>= 1
		k <<= 1
	}
	res := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		res = 2*res + tmp[i]
	}
	return res
}

//带点脑中，从后向前
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ret := 1<<maximumBit - 1
	res := make([]int, n)
	//最后一个肯定只有第一个数字在异或，其他都是逐步和前一个结果异或
	res[0] = nums[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] ^ nums[i]
	}
	//跟全1异或就是得出结果
	for i := 0; i < n; i++ {
		res[i] ^= ret
	}
	//记得翻转结果
	return reverse(res)
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

//从后往前，不翻转
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ret := 1<<maximumBit - 1
	res := make([]int, n)
	res[n-1] = nums[0]

	for i := n - 2; i >= 0; i-- {
		res[i] = res[i+1] ^ nums[n-i-1]
	}

	for i := 0; i < n; i++ {
		res[i] ^= ret
	}
	return res
}

func main() {
	nums := []int{0, 1, 1, 3}
	maximumBit := 2
	fmt.Println(getMaximumXor(nums, maximumBit))
}

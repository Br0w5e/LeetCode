package main

//476. 数字的补数
//1009. 十进制整数的反码
//N的二进制为xxxx，则结果为1111-xxxx
//与bitwiseComplement相同
func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	i := 1
	for i <= num {
		i <<= 1
	}
	return i-1 - num
}

//模拟
func findComplement1(num int) int {
	if num == 0 {
		return 1
	}
	nums := make([]int, 0)
	for num != 0 {
		nums = append(nums, num%2)
		num /= 2
	}
	res := 0
	for i := len(nums)-1; i >= 0; i-- {
		nums[i] ^= 1
		res = res*2 + nums[i]
	}
	return res
}
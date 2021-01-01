package main

import "fmt"
//暴力--> 超时
func getMaxLen(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[i] = 1
		} else if nums[i] < 0 {
			nums[i] = -1
		} else {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if judge(nums[i:]) > max {
			max = judge(nums[i:])
		}
	}
	return max
}

func judge(nums []int) int {
	mul := 1
	l := 0
	for i := 0; i < len(nums); i++ {
		mul *= nums[i]
		if mul > 0 {
			l = i+1
		}
	}
	return l
}

//递推
//f以正数结尾的最大长度，
//g以负数结尾的最大长度
func getMaxLen(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	g := make([]int, n)
	if nums[0] > 0 {
		f[0] = 1
	} else if nums[0] < 0{
		g[0] = 1
	}

	res := f[0]

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			f[i] = f[i-1]+1
			if g[i-1] != 0 {
				g[i] = g[i-1]+1
			}
		} else if nums[i] < 0 {
			if g[i-1] != 0 {
				f[i] = g[i-1] + 1
			}
			g[i] = f[i-1] + 1
		}
		res = max(res, f[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	nums := []int{70,-18,75,-72,-69,-84,64,-65,0,-82,62,54,-63,-85,53,-60,-59,29,32,59,-54,-29,-45,0,-10,22,42,-37,-16,0,-7,-76,-34,37,-10,2,-59,-24,85,45,-81,56,86}
	fmt.Println(getMaxLen(nums))
}

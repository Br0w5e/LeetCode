package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		//注意切片的传递内容
		tmp := make([]int, r[i]-l[i]+1)
		copy(tmp, nums[l[i]:r[i]+1])
		if judge(tmp) {
			res[i] = true
		}
	}
	return res
}

func judge(nums []int) bool{
	sort.Ints(nums)
	sub := nums[1]-nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != sub {
			return false
		}
	}
	return true
}

func main()  {
	nums := []int{4,6,5,9,3,7}
	l := []int{0,0,2}
	r := []int{2,3,5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}

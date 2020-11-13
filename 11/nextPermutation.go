package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	//寻找A[i]<A[j]，寻找第一个不是降序的位置
	//降序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	//不是最后一个
	if i >= 0 {
		//在[j, end]中寻找第一个大于nums[i]的
		for nums[i] >= nums[k] {
			k--
		}
		//交换nums[i], nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}
	//经过上面一步nums[j:]必然是降序，我们需要将其翻转使他成为升序
	//保证j到end是升序
	reverse(nums[j:])
	return
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
	return
}

func main()  {
	nums := []int{1,2,3}
	fmt.Println(nextPermutation(nums))
}

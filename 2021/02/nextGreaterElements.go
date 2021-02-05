package main

import "fmt"

//503. 下一个更大元素 II

//倍增数组，然后判断
func nextGreaterElements(nums []int) []int {
	res := make([]int, 0)
	l := len(nums)
	nums = append(nums, nums...)
	for i := 0; i < l; i++ {
		flag := 0
		for j := i; j < l+i; j++ {
			if nums[j] > nums[i] {
				res = append(res, nums[j])
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, -1)
		}
	}
	return res
}

//取模
func nextGreaterElements1(nums []int) []int {
	res := make([]int, 0)
	l := len(nums)
	for i := 0; i < l; i++ {
		flag := 0
		for j := i; j < l+i; j++ {
			if nums[j%l] > nums[i] {
				res = append(res, nums[j%l])
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, -1)
		}
	}
	return res
}

//单调栈
//存在下一个最大元素的数字应该被出栈入栈两次，否则就只能入栈一次
func nextGreaterElements2(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = -1
	}
	// 存储元素索引i
	stack := make([]int, 0)
	for i := 0; i < l*2; i++ {
		//当前元素有没有可能成为最大元素
		num := nums[i%l]
		// 对于当前元素，弹出栈中小于当前元素的元素，这些被弹出的元素的"下一个更大元素"就是当前元素
		// 一个元素只有入栈后再被弹出，才能得到该元素的"下一个更大元素"，否则无"下一个更大元素"，res[i]默认为-1
		// 例如，对于最大的元素,其一直留在栈中无法被弹出,因此其"下一个更大元素"默认为-
		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		//nums 中的每个元素都应该入栈，每个出入栈两次了
		stack = append(stack, i%l)
		//if i < l {
		//	stack = append(stack, i)
		//}
	}
	return res
}

func main()  {
	nums := []int{1,2,1}
	fmt.Println(nextGreaterElements2(nums))
}
package main

import "fmt"

ffunc minSubArrayLen(s int, nums []int) int {
if len(nums) == 0 {
return 0
}
res := make([]int, len(nums))
for i := 0; i < len(nums); i++ {
res[i] = getSum(nums, s, i)
}
rst := res[0]
for _, num := range res {
if num > 0 && num < rst {
rst = num
}
}
if rst < 0 {
return 0
}
return rst
}


func getSum(nums []int, s int, start int) int {
	sum := 0
	for i := start; i < len(nums); i++ {
		if sum < s && sum+nums[i] < s {
			sum += nums[i]
			continue
		} else if sum < s && sum+nums[i] >= s {
			return i-start+1
		} else {
			return i-start
		}
	}
	return -1
}

//规定复杂度里边
func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	sum := 0 // 和
	i := 0 //左指针
	j := 0 //右指针
	prevJ := -1 //上一次j指针，若发生了变动说明本次j可以计入和中
	for i < len(nums) && j < len(nums) {
		if j != prevJ { //j发生了变动，记录当前和
			sum += nums[j]
		}
		prevJ = j
		if sum >= s { //大于s，向右移动左指针
			if j - i + 1 < res { //记录结果
				res = j - i + 1
			}
			sum -= nums[i]
			if i == j { //如果i、j处于同一个位置，j也要向右移动
				i++
				j++
			} else {
				i++
			}
		} else { //小于s，继续向右移动j指针
			j++
		}
	}
	if res <= len(nums) {
		return res
	}
	return 0
}

func main()  {
	s := 3
	nums := []int{1,1}
	fmt.Println(minSubArrayLen(s, nums))
}
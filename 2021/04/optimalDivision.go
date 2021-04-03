package main


//553. 最优除法

//给数组加括号使得除法结果最大
//方法（第一个数字单独列出来，后边的所有数字依次做除法即可）
import "strconv"

func optimalDivision(nums []int) string {
	//长度为1和2需要单独处理
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	//一般情况
	res := ""
	for i, num := range nums {
		if i == 0 {
			res += strconv.Itoa(num) + "/("
		} else if i == len(nums)-1 {
			res += strconv.Itoa(num) + ")"
		} else {
			res += strconv.Itoa(num) + "/"
		}
	}
	return res
}

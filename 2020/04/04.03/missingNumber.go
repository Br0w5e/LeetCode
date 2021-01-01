//寻找缺失的数字
func missingNumber(nums []int) int {
	res := 0 
	for i, num := range nums {
		res += (i+1-num)
	}
	return res
}
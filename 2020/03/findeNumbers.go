//寻找数组中数字长度为偶数的 数字个数
//strconv.Itoa(v) 函数将 整数转化为对应的字符串
func findeNumbers(nums []int) int {
	res := 0
	for _, num := range nums{
		if getLen(num)%2 == 0{
			res++
		}
	}
	return res
}
func getLen(n int) int {
    switch {
        case n > -1 && n < 10:
            return 1
        case n > 9 && n < 100:
            return 2
        case n > 99 && n < 1000:
            return 3
        case n > 999 && n < 10000:
            return 4
        case n > 9999 && n < 100000:
            return 5
        }
    return 0
}
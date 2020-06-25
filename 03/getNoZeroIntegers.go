// 返回无0整数组合
func getNoZeroIntegers(n int) []int {
	left, right := 1, n-1
	for left <= right {
		strLeft := strconv.Itoa(left)
		strRight := strconv.Itoa(right)
		if strings.Contains(strLeft, "0") || strings.Contains(strRight, "0") {
			left++
			right--
		} else {
			return []int{left, right}
		}
	}
	return nil
}
//判断数组中每个数字出现次数是否是独一无二的
//一个map统计数字出现次数
//一个map统计次数出现的次数
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, num := range arr{
		m[num]++
	}
	times := make(map[int]int)
	for _, v := range m {
		times[v]++
		if times[v] > 1 {
			return false
		}
	}
	return true
}
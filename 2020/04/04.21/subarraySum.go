//和为K的子数列个数
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	s := 0
	// hash的意义在哪?差值出现的次数
	hash := make(map[int]int)
	// 初始化，差值为0，每个元素出现一次
	hash[0] = 1
	for _, v := range nums {
		s += v
		count += hash[s-k]
		hash[s]++
	}
	return count
}
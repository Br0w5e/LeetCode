//和为K的子数列个数
package  main

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
//暴力超时
func subarraySum2(nums []int, k int) int {
	n, cnt := len(nums), 0
	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {
			sum := 0
			for i := left; i <= right; i++ {
				sum += nums[i]
			}
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
//暴力 优化
func subarraySum3(nums []int, k int) int {
	n, cnt := len(nums), 0
	for left := 0; left < n; left++ {
		sum := 0
		for right := left; right < n; right++ {
			sum += nums[right]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

//最优子数组
func numberOfSubarrays(nums []int, k int) int {
	cnt := make([]int, len(nums)+1)
	res, odd := 0, 0
	cnt[0] = 1
	for _, num := range nums {
        //奇数
		if num&1 == 1 {
			odd++
		}
		cnt[odd]++
		if t := odd-k; t >= 0 {
			res += cnt[t]
		}
	}
	return res
}
package main

//1365. 有多少小于当前数字的数字
//暴力模拟
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		res = append(res, count(nums, num))
	}
	return res
}

func count(nums []int, target int) int {
	res := 0
	for _, num := range nums {
		if num < target {
			res++
		}
	}
	return res
}

//使用数组记录
func smallerNumbersThanCurrent1(nums []int) []int {
	cnt := make([]int, 101)
	for _, num := range nums {
		cnt[num]++
	}
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}
	res := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			res[i] = cnt[v-1]
		}
	}
	return res
}
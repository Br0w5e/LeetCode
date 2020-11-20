package main

//LCP 12. 小张刷题计划

//翻译题目：给定一个数组，将其划分成 MM 份，使得每份元素之和最大值最小，每份可以任意减去其中一个元素。
func minTime(time []int, m int) int {
	if m >= len(time) {
		return 0
	}
	left, right := time[0], 0
	for _, v := range time {
		right += v
		left = min(left, v)
	}
	ans := right
	for left <= right {
		mid := left + (right-left)/2
		sum, cnt, mx := 0, 1, 0
		for _, v := range time {
			mx = max(mx, v)
			if sum+v-mx > mid {
				cnt++
				sum, mx = v, v
			} else {
				sum += v
			}
		}
		if cnt <= m {
			ans = min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

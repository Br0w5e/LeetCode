package main
//1423. 可获得的最大点数

//滑动窗口
//从前往后滑
func maxScore(cardPoints []int, k int) int {
	res := 0
	l := len(cardPoints)
	for _, num := range cardPoints[:k] {
		res += num
	}
	sum := res
	for i := 1; i <= k; i++ {
		sum += (cardPoints[l-i]-cardPoints[k-i])
		if sum > res {
			res = sum
		}
	}
	return res
}

//从后往前滑
func maxScore(cardPoints []int, k int) int {
	res := 0
	l := len(cardPoints)
	for _, num := range cardPoints[l-k:] {
		res += num
	}
	sum := res
	for i := l-k; i < l; i++ {
		sum += (cardPoints[(i+k)%l]-cardPoints[i%l])
		if sum > res {
			res = sum
		}
	}
	return res
}
// [1,2,3,4,5,6,1, 1,2,3,4,5,6,1]

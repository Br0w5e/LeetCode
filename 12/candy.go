package main


//135. 分发糖果

func candy(ratings []int) int {
	res := 0
	n := len(ratings)
	//存储从左向右遍历递增结果
	left := make([]int, n)
	for  i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1]+1
		} else {
			left[i] = 1
		}
	}
	//从右遍历
	right := 0
	for i := n-1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		//当前位置的最大值应该等于右边最大的和左边最大的，取最大
		res += max(left[i], right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
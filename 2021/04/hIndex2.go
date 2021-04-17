package main


//275. H 指数 II

//参考 274. H 指数
//模拟，搞清思路，直接h
func hIndex(citations []int) int {
	n := len(citations)
	for h := len(citations); h > 0; h-- {
		if citations[n-h] >= h {
			return h
		}
	}
	return 0
}

//数量
func hIndex(citations []int) int {
	n := len(citations)
	for i, c := range citations {
		h := n-i
		if c >= h {
			return h
		}
	}
	return 0
}

//二分
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left) >> 1
		//大于的数量等于citations[mid]
		if citations[mid] == n-mid {
			return n-mid
			//大于的数量大于citations[mid]
		} else if citations[mid] < n-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n - left
}
package main

//986. 区间列表的交集
func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		l, r := max(A[i][0], B[j][0]), min(A[i][1], B[j][1])
		if l <= r {
			res = append(res, []int{l, r})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
package main

//763. 划分字母区间
//使用map进行标记
func partitionLabels(S string) []int {
	res := make([]int, 0)
	m := make(map[byte]int)
	//记录每个元素出现的最后一个位置
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}

	start, end := 0, 0
	for i := 0; i < len(S); i++ {
		//取得该元素出现的最后一个位置
		end = max(end, m[S[i]])
		//该位置是最后一个位置
		if i == end {
			res = append(res, end-start+1)
			start = i+1
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
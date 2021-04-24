package main

//1048. 最长字符串链
//先排序，后dp+mao
import "sort"

func longestStrChain(words []string) int {
	// 先排序
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	// mmp用于记录每个word 的最长链
	mmp := make(map[string]int)
	// 遍历words 依次计算最长链
	for _, v := range words {
		max := 0
		for i := 0; i < len(v); i++ {
			// max 取值 max和 mmp中少了个字符+1的大值
			max = Max(max, mmp[v[0:i]+v[i+1:]]+1)
		}
		// 得到v的最长链
		mmp[v] = max
	}
	// 看map里谁最大
	max := 0
	for e := range mmp {
		max = Max(max, mmp[e])
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

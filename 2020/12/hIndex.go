package main

//274. H 指数
import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	for i := 0; i < n; i++ {
		h := n-i
		if h <= citations[i] {
			return h
		}
	}
	return 0
}


func hIndex2(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	for h := n; h > 0; h-- {
		//至少有h篇引用都大于h
		if citations[n-h] >= h {
			return h
		}
	}
	return 0
}
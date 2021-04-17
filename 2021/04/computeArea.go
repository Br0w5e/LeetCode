package main

//223. 矩形面积

//弄清情况分类，慢慢编写
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	//不重叠情况
	if D <= F || C <= E || B >= H || A >= G {
		return (D-B)*(C-A) + (H-F)*(G-E)
	}
	//重叠情况，总面积-重叠面积
	left := max(A, E)
	right := min(C, G)
	up := min(D, H)
	down := max(B, F)
	return (D-B)*(C-A) + (H-F)*(G-E) - (up-down)*(right-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

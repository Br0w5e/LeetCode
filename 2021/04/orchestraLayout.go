package main

//LCP 29. 乐团站位

//数学公式
func orchestraLayout(num int, xPos int, yPos int) int {
	x := xPos
	y := yPos
	n := num
	if x <= y {
		k := min(x, n-1-y)
		return (4*k*(n-k)+1+(x+y-k*2)-1)%9 + 1
	}
	k := min(y, n-1-x) + 1
	return (4*k*(n-k)+1-(x+y-(k-1)*2)-1)%9 + 1

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

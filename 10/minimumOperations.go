package main

import "math"

//LCP 19. 秋叶收藏集

const inf = math.MaxInt32 // 或 math.MaxInt64

func minimumOperations(leaves string) int {
	n := len(leaves)
	g := -1
	if leaves[0] == 'y' {
		g = 1
	}
	gmin := g
	ans := inf
	for i := 1; i < n; i++ {
		isYellow := boolToInt(leaves[i] == 'y')
		g += 2*isYellow - 1
		if i != n-1 {
			ans = min(ans, gmin-g)
		}
		gmin = min(gmin, g)
	}
	return ans + (g+n)/2
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

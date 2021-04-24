package main

//789. 逃脱阻碍者
import "math"

//终点等待即可
func escapeGhosts(ghosts [][]int, target []int) bool {
	dis := math.MaxInt32
	for _, ghost := range ghosts {
		if abs(ghost[0]-target[0])+abs(ghost[1]-target[1]) < dis {
			dis = abs(ghost[0]-target[0]) + abs(ghost[1]-target[1])
		}
	}
	return dis > abs(target[0])+abs(target[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

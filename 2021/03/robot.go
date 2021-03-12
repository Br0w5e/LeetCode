package main

import "strings"

//LCP 03. 机器人大冒险

//暴力模拟--->超时
func robot(command string, obstacles [][]int, x int, y int) bool {
	dx, dy := 0, 0
	for dx <= x && dy <= y {
		for i := 0; i < len(command) && dx <= x && dy <= y; i++ {
			if command[i] == 'U' {
				dy++
			} else {
				dx++
			}
			for _, obstacle := range obstacles {
				if obstacle[0] == dx && obstacle[1] == dy {
					return false
				}
			}

			if dx == x && dy == y {
				return true
			}
		}
	}
	return false
}

//根据R和U的数量判断某个点在不在路线上就好了
func robot(command string, obstacles [][]int, x int, y int) bool {
	// 如果目标点不在路径上，返回失败
	if !isOnThePath(command, x, y) {
		return false
	}
	for _, o := range obstacles {
		// 判断有效的故障点是否在路径上（故障的步数大于等于目标的点，视为无效故障）
		if (x+y > o[0]+o[1]) && isOnThePath(command, o[0], o[1]) {
			return false
		}
	}
	return true
}

func isOnThePath(command string, x int, y int) bool {
	uNum := strings.Count(command, "U")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "U")
	rNum := strings.Count(command, "R")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "R")
	if uNum == y && rNum == x {
		return true
	}
	return false
}

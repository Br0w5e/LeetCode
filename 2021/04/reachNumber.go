package main


//754. 到达终点数字


func reachNumber(target int) int {
	target = abs(target)
	i := 0
	for target > 0 {
		i++
		target = target - i
	}
	if target == 0 || target%2 == 0 {
		return i
	} else {
		return i + i%2 + 1
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

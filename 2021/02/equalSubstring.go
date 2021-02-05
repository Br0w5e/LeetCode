package main
//1208. 尽可能使字符串相等

//滑动窗口
//暴力滑动
func equalSubstring(s string, t string, maxCost int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		cost, l := maxCost, 0
		for l+i < len(s) && cost >= 0 {
			cost -= abs(int(s[l+i]) - int(t[l+i]))
			l++
		}
		if l > res && cost >= 0 {
			res = l
		} else if l > res && cost < 0 {
			res = l-1
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}


//仿照424 滑动
func equalSubstring(s string, t string, maxCost int) int {
	l, r := 0, 0
	cost := 0
	for r < len(s) {
		cost += abs(int(s[r])-int(t[r]))
		r++
		if cost > maxCost {
			cost -= abs(int(s[l])-int(t[l]))
			l++
		}
	}
	return r-l
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
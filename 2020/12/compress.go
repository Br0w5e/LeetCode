package main

//443. 压缩字符串

//双指针
import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	if n == 0 || n == 1 {
		return n
	}
	var slow, fast, cnt, k int
	for fast <= n {
		if fast != n && chars[slow] == chars[fast] {
			cnt++
		} else {
			chars[k] = chars[slow]
			k++
			if cnt > 1 {
				tmp := []byte(strconv.Itoa(cnt))
				for j, _ := range tmp {
					chars[k] = tmp[j]
					k++
				}
			}
			slow = fast
			cnt = 1
		}
		fast++
	}
	return k
}

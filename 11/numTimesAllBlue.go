package main


//1375. 灯泡开关 III

//找截至目前亮着的灯的编号最大值：k次操作一定会使得k盏灯亮，灯变蓝等价于1~k的所有灯都亮。
func numTimesAllBlue(light []int) int {
	cur, res := 0, 0
	for i, v := range light {
		if cur < v {
			cur = v
		}
		if cur == i+1 {
			res++
		}
	}
	return res
}
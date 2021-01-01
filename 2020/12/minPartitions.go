package main

//5626. 十-二进制数的最少数目

//统计最大的
func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		if int(n[i] - '0') > res {
			res = int(n[i] - '0')
		}
	}
	return res
}

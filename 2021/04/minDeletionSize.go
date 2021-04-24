package main

//944. 删列造序
//按列遍历即可
func minDeletionSize(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				res++
				break
			}
		}
	}
	return res
}

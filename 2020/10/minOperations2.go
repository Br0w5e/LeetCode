package main
//1598. 文件夹操作日志搜集器
func minOperations(logs []string) int {
	res := 0
	for _, s := range logs {
		if s == "../" {
			//到达根目录时不能继续返回
			if res > 0{
				res--
			}
		} else if s == "./" {
			continue
		} else {
			res++
		}
	}
	return res
}
package main

//371. 两整数之和

func getSum(a int, b int) int {
	return a+b
}

//参考这个
//https://leetcode-cn.com/problems/sum-of-two-integers/solution/wei-yun-suan-xiang-jie-yi-ji-zai-python-zhong-xu-y/
func getSum2(a int, b int) int {
	//进位为0的时候退出
	for b != 0 {
		//无进位和结果
		xor := a ^ b
		//进位
		b = (a & b) << 1
		a = xor
	}
	return a
}
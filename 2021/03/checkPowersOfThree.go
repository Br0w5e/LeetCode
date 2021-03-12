package main

//5681. 判断一个数字是否可以表示成三的幂的和

//在数字的三进制中不能出现2，只能出现0和1
func checkPowersOfThree(n int) bool {
	for n != 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}
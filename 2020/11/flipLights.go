package main

//672. 灯泡开关 Ⅱ

/*
所有的操作进行偶数次是会抵消的，那么所有操作只存在0,1两种，即无效果和有效果；
且各一次操作2,3等效为一次操作1；
画一个类似真值表，可以推出  n>=3且m>=3时,结果只会是8；
接下来考虑个别情况即可。
*/

func flipLights(n int, m int) int {
	if n > 3 {
		n = 3
	}
	if m == 0 {
		return 1
	} else if m == 1 {
		return []int{2, 3, 4}[n-1]
	} else if m == 2 {
		return []int{2, 4, 7}[n-1]
	}
	return []int{2, 4, 8}[n-1]
}

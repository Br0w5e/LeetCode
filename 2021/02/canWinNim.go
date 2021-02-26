package main

//292. Nim 游戏

//先下手的可以控制剩余的石子是4的倍数
func canWinNim(n int) bool {
	//return !(n%4 == 0)
	return n%4 != 0
}

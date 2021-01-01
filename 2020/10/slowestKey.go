package main

//5546. 按键持续时间最长的键
func slowestKey(releaseTimes []int, keysPressed string) byte {
	//存储结果
	res := keysPressed[0]
	//存储最大的下标
	max := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		//当前按键时长
		tmp := releaseTimes[i]-releaseTimes[i-1]
		if tmp > max {
			max = tmp
			res = keysPressed[i]
		} else if tmp == max && keysPressed[i] > res{
			res = keysPressed[i]
		}
	}
	return res
}

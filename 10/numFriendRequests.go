package main
//825. 适龄的朋友

//暴力破解 --> 超时   10^7
func numFriendRequests(ages []int) int {
	res := 0
	for i := 0; i < len(ages); i++ {
		for j := 0; j < len(ages); j++ {
			if ages[i] <= ages[j]/2+7 || ages[i] > ages[j] || (ages[i] > 100 && ages[j] < 100) || i == j {
				continue
			} else {
				res++
			}
		}
	}
	return res
}

//
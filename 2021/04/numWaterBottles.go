package main

//1518. 换酒问题

//不存在借酒瓶的现象
func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		res += numBottles/numExchange
		numBottles = numBottles%numExchange + numBottles/numExchange
	}
	return res
}
// 15+3

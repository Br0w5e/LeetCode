package main

//888. 公平的糖果棒交换

//哈希表
func fairCandySwap(A []int, B []int) []int {
	mA := make(map[int]bool, 0)
	mB := make(map[int]bool, 0)
	diff := 0
	for _, a := range A {
		diff += a
		mA[a] = true
	}
	for _, b := range B {
		diff -= b
		mB[b] = true
	}

	diff /= 2
	for k, _ := range mA {
		if mB[k-diff] == true {
			return []int {k, k-diff}
		}
	}
	return []int{}
}

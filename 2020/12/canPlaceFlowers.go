package main

//605. 种花问题

func canPlaceFlowers(flowerbed []int, n int) bool {
	sum := 0
	black := 1
	for _, num := range flowerbed {
		if num == 0 {
			black++
		} else {
			sum += (black-1)/2
			black = 0
		}
	}
	sum += black/2
	return sum >= n
}

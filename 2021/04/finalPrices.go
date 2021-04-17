package main


//1475. 商品折扣后的最终价格

//遍历
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i+1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

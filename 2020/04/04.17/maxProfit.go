//买卖股票的最大收益
//一次遍历
func maxProfit(prices []int) int {
    maxprofit := 0 
    for i := 1; i < len(prices); i++{
        if prices[i] > prices[i-1] {
            maxprofit += (prices[i]-prices[i-1])
        }
    }
    return maxprofit
}
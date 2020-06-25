//计算股票的最大收益
//方法一： 逐个计算，最后比较会超出内存范围，进行优化
func maxProfit(prices []int) int {
    res := make([]int, 0)
    n := len(prices)
    if n == 0  || n == 1{
        return 0 
    }
    for i := 0; i < n-1; i++{
        for j := i+1; j < n; j++ {
            res = append(res, prices[j]-prices[i])
        }
    }
    max := res[getMax(res)]
    if max < 0{
        return 0
    }
    return max
}

func getMax(nums []int) int {
    left := 0
    right := len(nums) - 1
    for left < right {
        if nums[left] < nums[right]{
            left++
        } else {
            right--
        }
    }
    return left
}


//动态规划
func maxProfit(prices []int) int {
    dp_0 := 0
    dp_1 := 0x7fffffff
    for i := 0; i < len(prices); i++{
        dp_0 = max(dp_0, prices[i]-dp_1)
        dp_1 = min(dp_1, prices[i])
    }
    return dp_0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

//另方法
func maxProfit(prices []int) int {
    price := 21000000
	maxprofit := -21000000
	if len(prices) == 0{
		return 0
	}
	for i := 0 ; i < len(prices) ; i ++ {
		//最小入价
		if prices[i] < price {
			price = prices[i]
		}
		//最大出价
		if prices[i] - price > maxprofit {
			maxprofit = prices[i] - price
		}
	}
	return maxprofit
}
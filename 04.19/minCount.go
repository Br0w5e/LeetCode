//拿硬币
//每次只能拿一个或者两个求拿完所有硬币的步伐
//暴力模拟
func minCount(coins []int) int {
    res := 0
    for i := 0; i < len(coins); i++{
        for coins[i] != 0 {
            if coins[i] >= 2 {
                coins[i] -= 2
            } else {
                coins[i]--
            }
            res++
        }
    }
    return res
}
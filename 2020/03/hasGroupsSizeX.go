// 卡牌分组
//判断卡牌能不能分成指定的分组
//思路 通过判断数字出现个数的公约数

func hasGroupsSizeX(deck []int) bool {
    max := getMax(deck)
    res := make([]int, max+1)
    for _, cnt := range deck {
        res[cnt]++
    }
    g := res[0]
    for i := 1; i < max+1; i++{
        g = gcd(g, res[i])
    }
    return g >= 2
}

func getMax(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        if nums[left] < nums[right]{
            left++
        } else{
            right--
        }
    }
    return nums[left]
}

func gcd(a, b int) int{
    if a < b{
        a, b = b, a
    }
    for b != 0{
        a, b = b, a%b
    }
    return a
}
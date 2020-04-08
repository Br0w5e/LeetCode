//转化996
//方法：先转化为数组再改变，最后转化回来。
func maximum69Number(num int) int {
    tmp := turnToStr(num)
    for i := len(tmp)-1; i >= 0; i--{
        if tmp[i] == 6{
            tmp[i] = 9
            break
        }
    }
    return turnTonNum(tmp) 
}

func turnToStr(num int) []int {
    res := make([]int, 0)
    for num != 0 {
        res = append(res, num%10)
        num /= 10
    }
    return res
}

func turnTonNum(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        res += int(math.Pow(10, float64(i))) * nums[i]
    }
    return res
}
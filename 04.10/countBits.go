//比特位计数，计算每一个小于等于k的每个数字中有多少个1
//方法1：朴素方法
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 0; i <= num; i++{
        res[i] = getOne(i)
    }
    return res
}

func getOne(num int) int {
    res := 0
    for num != 0 {
        res += num%2
        num /= 2
    }
    return res
}

//方法2：i 和 i/2 要么1的个数相同（偶数）， 要么多一个1（奇数）
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++ { 
        // i 和 i/2 要么1的个数相同（偶数）， 要么多一个1（奇数）
        //res[i] = res[i>>1] + (i&1)
        if i%2 == 0 {
            res[i] = res[i/2]
        } else {
            res[i] = res[i/2]+1
        }
    }
    return res
}
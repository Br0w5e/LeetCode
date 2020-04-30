//面试题 01.02. 判定是否互为字符重排
//方法：字典或者数组记录单个字母在s1中出现次数，并在s2中减去，然后判断字典所有键的值是不是0
func CheckPermutation(s1 string, s2 string) bool {
    res := make(map[rune]int)
    for _, s := range s1 {
        res[s]++
    } 
    for _, s := range s2 {
        res[s]--
        if res[s] < 0 {
            return false
        }
    }
    for _, value := range res {
        if value != 0{
            return false
        }
    }
    return true
}
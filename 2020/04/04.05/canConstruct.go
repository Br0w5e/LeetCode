//判断字符串可不可以构成指定的回文串。
//原理性质
func canConstruct(s string, k int) bool {
    table := make([]int, 26)
    for _, value := range s {
        table[value-'a']++
    }
    res := 0
    for i := 0; i < 26; i++{
        if table[i]%2 == 1 {
            res ++
        }
    }
    return res <= k && k <= len(s)
}
//重新格式化字符串
//思路：先统计数字和字母个数，再进行排列即可，两个指针
func reformat(s string) string {
    numNum, chNum := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            numNum++
        } else {
            chNum++
        }
    }
    if chNum-numNum > 1 || chNum-numNum < -1 {
        return ""
    }
    res :=[]byte(s)
    if chNum > numNum {
        for i := 0; i < len(s); i++{
            for j := len(s)-1; j >= 0; j-- {
                if i%2 == 0 && j%2==1 && res[i] >= '0' && res[i] <= '9' && res[j] >= 'a' && res[j] <= 'z' {
                    res[i], res[j] = res[j], res[i]
                }
            }
        }   
    } else {
        for i := 0; i < len(s); i++{
            for j := len(s)-1; j >= 0; j-- {
                if i%2 == 0 && j%2==1 && res[i] >= 'a' && s[i] <= 'z' && res[j] >= '0' && res[j] <= '9' {
                    res[i], res[j] = res[j], res[i]
                }
            }
        }   
    }
    return string(res)
}
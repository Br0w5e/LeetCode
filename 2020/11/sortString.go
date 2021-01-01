package main

//1370. 上升下降字符串

//将字符串中的字母按照：升序走一趟，降序走一趟的顺序进行排列
//思路：统计每个字母出现的次数；然后按照升序走一趟降序走一趟的思路，知道数组全部为0
func sortString(s string) string {
    cnt := make([]int, 26)
    for _, val := range s {
        cnt[val-'a']++
    }
    res := ""
    flag := 0
    for !allZero(cnt){
		//升序
        if flag%2 == 0 {
            for i := 0; i < 26; i++{
                if cnt[i] != 0 {
                    res += string('a'+i)
                    cnt[i]--                    
                }
			}
		//降序
        } else {
            for i := 25; i >= 0; i--{
                if cnt[i] != 0 {
                    res += string('a'+i)
                    cnt[i]--       
                }
            }
        }
        flag++
    } 
    return res
}
func allZero(nums []int) bool {
    for _, num := range nums {
        if num != 0 {
            return false
        }
    }
    return true
}
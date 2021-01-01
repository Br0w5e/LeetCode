//判断出勤情况
//方法一：两次遍历
func checkRecord(s string) bool {
    absent := 1
    for _, letter := range s {
        if letter == 'A'{
            absent--
        }
    }
    for i := 0; i < len(s)-2; i++{
        if s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L'{
            return false
        }
    }
    if absent >= 0 {
        return true
    }
    return false
}

//方法二： 一次遍历
func checkRecord(s string) bool {
    absent := 0
    late := 0
    for _,v := range s {
        if v == 'A' {
            absent++
            late = 0
        } else if v == 'L' {
            late++
        } else {
            late = 0
        }
        if absent > 1 || late > 2 {
            return false
        }
    }
    return true
}
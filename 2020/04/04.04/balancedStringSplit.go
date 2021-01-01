//分割平衡字符串
//将一平衡字符串，分割为更多的平衡字符串。
//方法：设置flag位，每当归零时候则分割出一个平衡字符串
func balancedStringSplit(s string) int {
    flag, res := 0, 0
    for _, value := range s {
        if value == 'L' {
            flag++
        } else {
            flag--
        }
        if flag == 0 {
            res++
        }
    }
    return res
}
//判断机器人能否回到原点
//向左走的步数应该和向右走的步数一样，向上走的步数应该和向下走的步数一样
//用单个flag标记或者多个flag标记
func judgeCircle(moves string) bool {
    flag := 0
    for _, move := range moves {
        switch move {
            case 'L': flag++
            case 'R': flag--
            case 'U': flag += 2
            case 'D': flag -= 2
        }
    }
    if flag == 0{
        return true
    } 
    return false
}
//两字符串中仅一个b比a多一个字母，其他字母都相同，求出该字母
//map
func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)
    for _, v := range t{
        m[byte(v)]++
    }
    for _, v := range s {
        m[byte(v)]--
    }
    for k, v := range m {
        if v == 1 {
            return k
        }
    }
    return 0
}
//异或
func findTheDifference(s string, t string) byte {
    var res byte
    for _, ele := range s {
        res ^= byte(ele)
    }
    for _, ele := range t {
        res ^= byte(ele)
    }
    return res
}
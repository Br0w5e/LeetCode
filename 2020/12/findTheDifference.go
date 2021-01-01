package main

//两字符串中仅一个b比a多一个字母，其他字母都相同，求出该字母
//map
func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
        m[t[i]]--
    }
    for k, v := range m {
        if v < 0 {
            return k
        }
    }
    return t[len(t)-1]
}

//异或
func findTheDifference1(s string, t string) byte {
    var res byte
    for i := 0; i < len(s); i++ {
        res ^= s[i]
    }
    for i := 0; i < len(t); i++ {
        res ^= t[i]
    }
    return res
}

//求和
func findTheDifference2(s string, t string) byte {
    sum := 0
    for i := 0; i < len(s); i++ {
        sum -= int(s[i]-'a')
        sum += int(t[i]-'a')
    }
    sum += int(t[len(t)-1]-'a')
    return byte(sum+'a')
}
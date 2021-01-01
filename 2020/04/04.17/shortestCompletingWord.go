//最短完整单词
//先写一个匹配函数
//匹配函数用map实现
func shortestCompletingWord(licensePlate string, words []string) string {
    licensePlate = strings.ToLower(licensePlate)
    var res string
    for _, word := range words {
        if match(licensePlate, word) {
            if res == "" || len(word) < len(res) {
                res = word
            }
        }
    }
    return res
}
func match(text string, word string) bool {
    m := make(map[byte]int)
    for _, ele := range text {
        if ele >= 'a' && ele <= 'z' {
            m[byte(ele)]++
        }
    }
    for _, ele := range word {
        if m[byte(ele)] > 0 {
            m[byte(ele)]--
        }
    }
    for _, v := range m {
        if v > 0 {
            return false
        }
    }
    return true
}
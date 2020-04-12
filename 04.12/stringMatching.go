//5380. 数组中的字符串匹配
//方法：写一个匹配函数 然后遍历两次，记得标记访问过的元素
func stringMatching(words []string) []string {
    res := make([]string, 0)
    m := make(map[string]bool)
    for i := 0; i < len(words); i++{
        for j := 0; j < len(words); j++{
            if i == j {
                continue
            }
            if !m[words[j]] && match(words[i], words[j]) {
                res = append(res, words[j])
                m[words[j]] = true
            }
        }
    }
    return res
}
func match(text string, target string) bool{
    if len(target) > len(text) {
        return false
    }
    if target == text {
        return true
    }
    for i := 0; i < len(text)-len(target)+1; i++{
        if target == text[i:len(target)+i] {
            return true
        }
    }
    return false
}

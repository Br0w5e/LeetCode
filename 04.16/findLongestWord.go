//通过删除字母匹配字典里的最长单词
//双指针，先判断是否包含，然后再比较长短
func findLongestWord(s string, d []string) string {
    var res string
    for _, word := range d {
        if compare(s, word) {
            if len(word) > len(res) || (len(word) == len(res) && strings.Compare(word, res) < 0) {
                res = word
            }
        }
    }
    return res
}

func compare(s string, word string) bool {
    sLen, wordLen := len(s), len(word)
    i, j := 0, 0
    for i < sLen && j < wordLen {
        if s[i] == word[j]{
            j++
        }
        i++
    }
    return j == wordLen
}
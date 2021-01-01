//了解掌握一下go中 string、byte、rune的区别
func reverseVowels(s string) string {
    sr := []rune(s)
    left, right := 0, len(s)-1
    for left < right{
        if !isVowel(sr[left]){
            left++
            continue
        }
        if !isVowel(sr[right]){
            right--
            continue
        }
        sr[left], sr[right] = sr[right], sr[left]
        left++
        right--
    }
    return string(sr)
}

func isVowel(ch rune) bool{
    ch = unicode.ToLower(ch)
    return ch == 'a' || ch == 'o' || ch == 'e' || ch == 'u' || ch == 'i'
}
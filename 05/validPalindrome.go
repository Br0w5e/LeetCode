//判断给定字符串，在最多删去一个字符的条件下是不是回文串
//思路遇到不相等的进行跳跃
package main
func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            break
        }
        left++
        right--
    }
    if left > right {
        return true
    }
    return help(s, left+1, right) || help(s, left, right-1)
}

func help(s string, left int, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
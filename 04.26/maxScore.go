//最大得分
func maxScore(s string) int {
    max := 0
    for i := 1; i < len(s); i++ {
        if cntZero(s[:i]) + cntOne(s[i:]) > max {
            max = cntZero(s[:i]) + cntOne(s[i:])
        }
    }
    return max
}

func cntZero(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '0' {
            res++
        }
    }
    return res
}

func cntOne(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            res++
        }
    }
    return res
}
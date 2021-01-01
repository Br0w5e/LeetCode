//字母易位次
//两个字符串
func isAnagram(s string, t string) bool {
    sArr := strings.Split(s, "")
    tArr := strings.Split(t, "")
    sort.Strings(sArr)
    sort.Strings(tArr)
    return strings.Join(sArr, "") == strings.Join(tArr, "")
}

//字符串数组中的易位次
func groupAnagrams(strs []string) [][]string {
    res := [][]string{}
    if len(strs) == 0 {
        return res
    }
    helpMap := make(map[string][]string)
    for i := 0; i < len(strs); i++{
        curStrArr := strings.Split(strs[i], "")
        sort.Strings(curStrArr)
        curStrIndex := strings.Join(curStrArr, "")
        helpMap[curStrIndex] = append(helpMap[curStrIndex], strs[i])
    }
    for _, value := range helpMap {
        res = append(res, value)
    }
    return res
}
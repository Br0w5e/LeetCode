//压缩编码/合并前缀
func minimumLengthEncoding(words []string) int {
	//建立映射表
	wmap := make(map[string]bool, len(words))
	for _, word := range words {
		wmap[word] = true
	}
	//删除重复元素
	for _, word := range words {
		for i := 1; i < len(word); i++{
			_, exist := wmap[word[i:]]
			if exist {
				delete(wmap, word[i:])
			}
		}
	}
	//返回结果
	res := 0
	for k, _ := range wmap{
		res += (len(k) + 1)
	}
	return res
}

//方案二
func minimumLengthEncoding(words []string) int {
    var setWord []string
    tmpMap := make(map[string]int, 0)
    for _, word := range words{
        if _, ok := tmpMap[word]; !ok {
            setWord = append(setWord, reverse(word))
            tmpMap[word] = 1
        }
    }
    sort.Strings(setWord)
    count := 0
    for i := 0; i < len(setWord)-1; i++{
        if !strings.HasPrefix(setWord[i+1], setWord[i]) {
            count += len(setWord[i]) + 1
        }
    }
    count += len(setWord[len(setWord)-1]) + 1
    return count
}

func reverse(s string) string {
    runes := []rune(s)
    left, right := 0, len(runes)-1
    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }
    return string(runes)
}


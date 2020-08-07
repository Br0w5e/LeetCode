package main

import "runtime/debug"

//暴力 方法 超时
func palindromePairs(words []string) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && judgePalindrome(words[i]+words[j]) {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func judgePalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}


func init() { debug.SetGCPercent(-1) }

func palindromePairs2(words []string) [][]int {
	strMap := make(map[string]int, len(words))
	for i, word := range words {
		strMap[word] = i
	}
	ret := make([][]int, 0)
	for i, word := range words {
		for k := 0; k < len(word); k++ {
			prefix, suffix := word[0:k], word[k:]
			//case1
			if isPali(prefix) {
				//前缀回文，后缀翻转后存在且不是自己
				//k=0时suffix取全字符串，覆盖case3
				//k最多取到len-1，后缀不为空，不用考虑case4
				if j, ok := strMap[reversePali(suffix)]; ok && i != j {
					ret = append(ret, []int{j, i})
				}
			}
			//case2
			if isPali(suffix) {
				//后缀回文，前缀翻转后存在且不是自己
				//k不能取到len，否则k=0和k=len重复，故prefix取不到全字符串，不用考虑case3，
				//前缀为空，翻转为空，空不能进行遍历，此时加上，覆盖case4
				if j, ok := strMap[reversePali(prefix)]; ok && i != j {
					ret = append(ret, []int{i, j})
					if k == 0 {
						ret = append(ret, []int{j, i})
					}
				}
			}
		}
	}
	return ret
}

//翻转字符串
func reversePali(word string) string {
	chars := []byte(word)
	for low, high := 0, len(chars)-1; low < high; low, high = low+1, high-1 {
		chars[low], chars[high] = chars[high], chars[low]
	}
	return string(chars)
}

//判断字符串是不是回文
func isPali(word string) bool {
	for low, high := 0, len(word)-1; low < high; low, high = low+1, high-1 {
		if word[low] != word[high] {
			return false
		}
	}
	return true
}
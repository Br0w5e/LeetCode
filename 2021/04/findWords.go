package main

import "strings"

//500. 键盘行
func findWords(words []string) []string {
	keyboard := map[int][]byte{
		1: []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		2: []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		3: []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}

	keyIndexs := map[byte]int{}

	for k, v := range keyboard {
		for _, x := range v {
			keyIndexs[x] = k
		}
	}

	result := []string{}

	for _, v := range words {
		value := strings.ToLower(v)

		firstIndex := keyIndexs[value[0]]

		flag := true

		for i := 1; i < len(value); i++ {
			if keyIndexs[value[i]] != firstIndex {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, v)
		}
	}

	return result
}

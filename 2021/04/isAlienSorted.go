package main

//953. 验证外星语词典

//先还原成正常排序
import "fmt"

func isAlienSorted(words []string, order string) bool {
	byteM := make(map[byte]byte)
	for i, ch := range order {
		byteM[byte(ch)] = byte(i + 'a')
	}
	wordsList := make([]string, 0)
	for _, word := range words {
		tmp := make([]byte, 0)
		for i := 0; i < len(word); i++ {
			tmp = append(tmp, byteM[word[i]])
		}
		wordsList = append(wordsList, string(tmp))
	}
	for i := 0; i < len(wordsList)-1; i++ {
		if wordsList[i] > wordsList[i+1] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"kuvp", "q"}
	order := "ngxlkthsjuoqcpavbfdermiywz"
	fmt.Println(isAlienSorted(words, order))
}

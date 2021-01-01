package main

//5557. 最大重复子字符串
import (
	"fmt"
)

func maxRepeating(sequence string, word string) int {
	tmp := 0
	if len(sequence) < len(word) || len(word) == 0{
		return 0
	}
	if sequence == word {
		return 1
	}
	max := 0
	for i := 0; i < len(sequence)-len(word)+1; {
		if sequence[i:i+len(word)] == word {
			tmp++
			i += len(word)
		} else {
			tmp = 0
			i++
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main()  {
	s := "aaa"
	w := "a"
	fmt.Println(maxRepeating(s, w))
}
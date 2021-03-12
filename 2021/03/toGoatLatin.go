package main

//824. 山羊拉丁文
import "strings"

func toGoatLatin(S string) string {
	//S := strings.ToLower(S)
	arr := strings.Split(S, " ")
	for i, s := range arr {
		sByte := []byte(s)
		if !(sByte[0] == 'a' || sByte[0] == 'e' || sByte[0] == 'i' || sByte[0] == 'o' || sByte[0] == 'u' || sByte[0] == 'A' || sByte[0] == 'E' || sByte[0] == 'I' || sByte[0] == 'O' || sByte[0] == 'U') {
			head := sByte[0]
			sByte = sByte[1:]
			sByte = append(sByte, head)
		}
		sByte = append(sByte, byte('m'))
		sByte = append(sByte, byte('a'))
		for j := 0; j < i+1; j++ {
			sByte = append(sByte, byte('a'))
		}
		arr[i] = string(sByte)
	}
	return strings.Join(arr, " ")
}
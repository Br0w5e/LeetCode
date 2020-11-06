package main

//5554. 能否连接形成数组
func canFormArray(arr []int, pieces [][]int) bool {
	i := 0
	for i < len(arr) {
		tmp := i
		for _, piece := range pieces {
			if piece[0] == arr[i] {
				for j := 0; i < len(arr) && j < len(piece); {
					if arr[i] == piece[j] {
						i++
						j++
					} else {
						return false
					}
					if i == len(arr) {
						return true
					}
				}
			}
		}
		if i == tmp {
			return i == len(arr)
		}
	}
	return i == len(arr)
}

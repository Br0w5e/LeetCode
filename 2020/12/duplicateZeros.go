package main

//1089. 复写零
//模拟
func duplicateZeros(arr []int)  {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i < len(arr)-1{
			copy(arr[i+1:], arr[i:])
			arr[i+1] = 0
			i++
		}
	}
}


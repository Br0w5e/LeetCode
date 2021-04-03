package main

//393. UTF-8 编码验证
//与操作
func validUtf8(data []int) bool {
	cnt := 0
	for _, v := range data {
		if cnt == 0 {
			if v&0b11111000 == 0b11110000 {
				cnt = 3
			} else if v&0b11110000 == 0b11100000 {
				cnt = 2
			} else if v&0b11100000 == 0b11000000 {
				cnt = 1
			} else if v&0b10000000 != 0 {
				return false
			}
		} else {
			if v&0b10000000 != 0b10000000 {
				return false
			}
			cnt--
		}
	}

	return cnt == 0
}
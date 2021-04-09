package main

//5705. 判断国际象棋棋盘中一个格子的颜色

//模拟
func squareIsWhite(coordinates string) bool {
	switch coordinates[0] {
	case 'a', 'c', 'e', 'g':
		if (coordinates[1]-'1')%2 == 0 {
			return false
		}
		return true
	case 'b', 'd', 'f', 'h':
		if (coordinates[1]-'1')%2 == 1 {
			return false
		}
		return true
	}
	return false
}

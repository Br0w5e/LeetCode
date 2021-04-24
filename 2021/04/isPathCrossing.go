package main

//1496. 判断路径是否相交

//map标记
func isPathCrossing(path string) bool {
	N, E := 0, 0
	m := make(map[int]bool)
	m[0] = true
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			N++
		case 'S':
			N--
		case 'E':
			E++
		case 'W':
			E--
		}
		if m[N*10000+E] {
			return true
		}
		m[N*10000+E] = true
	}
	return false
}

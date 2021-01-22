// 1232. 缀点成线
//检验所有点是不是在同一直线上
//方法：斜率,用乘法别用除法
func checkStraightLine(c [][]int) bool {
	x, y := c[1][0]-c[0][0], c[1][1]-c[0][1]
	for i := 2; i < len(c); i++ {
		xi, yi := c[i][0]-c[0][0], c[i][1]-c[0][1]
		if x*yi != xi*y {
			return false
		}
	}
	return true
}
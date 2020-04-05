//判断矩形和圆的位置关系
//思路：初中知识，判断矩形边上的点与圆心距离和半径的关系
func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	for i := x1; i <= x2; i++{
		for j := y1; j <= y2; j++{
			if (i-x_center)*(i-x_center) + (j-y_center)*(j-y_center) <=  radius*radius {
				return true
			}
		}
	}
	return false
}
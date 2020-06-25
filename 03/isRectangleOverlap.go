//判断矩形是否重叠
//思路：通过坐标差进行判断
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    return (rec2[0] - rec1[2]) * (rec2[2] - rec1[0]) < 0 && (rec2[1] - rec1[3]) * (rec2[3] - rec1[1]) < 0
}
//思路：通过坐标关系进行判断
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[0] >= rec2[2] || rec1[2]  <= rec2[0] || rec1[1] >= rec2[3] || rec1[3] <= rec2[1]{
	  return false
	}
	  return true
  }
//锯齿形遍历 参考levelOrder.go
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		tmp := []*TreeNode{}
		res = append(res, make([]int, 0))
		for _, v := range queue {
			if level % 2 == 0{
				res[level] = append(res[level], v.Val)
			} else {
				s := []int{v.Val}
				res[level] = append(s, res[level]...)
			}
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
			
		}
		queue = tmp
		level++
	}
	return res
}
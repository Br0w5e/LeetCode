package main

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	//存储每一层所有节点
	queue := make([]*TreeNode, 0)
	//第一层只有跟节点
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		var sum float64
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, sum/float64(l))
	}
	return res
}

package main

//429. N 叉树的层序遍历
/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
 */

//bfs
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for size > 0 {
			size--
			p := queue[0]
			queue = queue[1:]
			tmp = append(tmp, p.Val)
			queue = append(queue, p.Children...)
		}
		res = append(res, tmp)
	}
	return res
}

//dfs
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *Node, depth int, res *[][]int)  {
	if root == nil {
		return
	}
	if len(*res) <= depth {
		(*res) = append((*res), make([]int, 0))
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	for _, ch := range root.Children {
		dfs(ch, depth+1, res)
	}
}


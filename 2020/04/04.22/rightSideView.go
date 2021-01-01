//二叉树的右视图
//层次遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //dfs右节点先序遍历
 func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res, 1)
	return res
}

func dfs(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if level > len(*res) {
		*res = append(*res, root.Val)
	}
	dfs(root.Right, res, level+1)
	dfs(root.Left, res, level+1)
}

//bfs右节点先序遍历
func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil{
        return res
    }
	queue := []*TreeNode{root}
	level := 0
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			if len(res) == level {
				res = append(res, queue[0].Val)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			queue = queue[1:]
		}
		level++
	}
	return res
}
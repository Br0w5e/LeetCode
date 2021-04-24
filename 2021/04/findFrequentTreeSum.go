package main

//508. 出现次数最多的子树元素和

//遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)
	queue := make([]*TreeNode, 0)
	travel(root, &queue)
	for i := 0; i < len(queue); i++ {
		sum := 0
		dfs(queue[i], &sum)
		m[sum]++
	}
	tmp := 0
	for _, v := range m {
		if v > tmp {
			tmp = v
		}
	}
	res := make([]int, 0)
	for k, v := range m {
		if v == tmp {
			res = append(res, k)
		}
	}
	return res

}

//计算和
func dfs(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res += root.Val
	dfs(root.Right, res)
}

//树遍历
func travel(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	travel(root.Left, res)
	*res = append(*res, root)
	travel(root.Right, res)
}

//求二叉搜索树指范围的和
//当前搜索值小于L时候搜索其右子树，当前搜索值大于于R时候搜索其左子树。
func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil {
        return 0
    }
    if root.Val < L {
        return rangeSumBST(root.Right, L, R)
    }
    if root.Val > R {
        return rangeSumBST(root.Left, L, R)
    }
    return root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}

//dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum int

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum = 0
	dfs(root, L, R)
	return sum
}

func dfs(node *TreeNode, L int, R int) {
	if node != nil {
		if node.Val >= L && node.Val <= R {
				sum += node.Val
				dfs(node.Left, L, R)
				dfs(node.Right, L, R)
			}
		if node.Val < L {
			dfs(node.Right, L, R)
		}
		if node.Val > R {
			dfs(node.Left, L, R)
		}
	} 
}
 
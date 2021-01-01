//判断二叉树是否对称
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSame(root.Left, root.Right)
}

func isSame(Left *TreeNode, Right *TreeNode) bool {
    if Left == nil &&  Right == nil {
        return true
    }
    if (Left != nil && Right == nil) || (Left == nil && Right != nil) {
        return false
    }
    return Left.Val == Right.Val && isSame(Left.Left, Right.Right) && isSame(Left.Right, Right.Left)
}

//深度优先遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root.Left, root.Right)
}
func dfs(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil || root2 == nil {
        return false
    }
    //这一步必须return false用不等于 不能用等于 return true
    if root1.Val != root2.Val {
        return false
    }
    return dfs(root1.Left, root2.Right) && dfs(root1.Right, root2.Left)
}
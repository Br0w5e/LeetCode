//判断一个树是否是平衡二叉树
//思路：逐个递归判断

//面试题 04.04. 检查平衡性
package main
func isBalanced(root *TreeNode) bool {
    if root == nil{
        return true
    }
    tp := depth(root.Right) - depth(root.Left)
    if tp == -1 || tp == 0 || tp == 1 {
        return isBalanced(root.Right) && isBalanced(root.Left)
    }
    return false
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(depth(root.Left), depth(root.Right))+1
}

func max(a, b int) int {
    if a > b{
        return a
    }
    return b
}
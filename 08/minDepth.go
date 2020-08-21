//寻找二叉树的最小深度
// 方法：递归深度优先。或者广度优先
package main
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var depth int
    trees := []*TreeNode{root}
    for len(trees) > 0 {
        var temp []*TreeNode
        for _,tree := range trees{
            if tree.Left == nil && tree.Right == nil {
                return depth + 1
            }
            if tree.Left != nil{
                temp = append(temp,tree.Left)
            }
            if tree.Right != nil{
                temp = append(temp,tree.Right)
            }
        }
        depth++
        trees = temp
    }
    return depth
}

//递归法
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := minDepth(root.Left)
    right := minDepth(root.Right)

    if left == 0 {
        return right+1
    }

    if right == 0 {
        return left+1
    }
    return min(left, right)+1
}

func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}
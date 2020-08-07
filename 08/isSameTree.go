//判断两个二叉树是否相同
//进行深度优先遍历即可
package main
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return dfs(p, q)
}

func dfs(root1 *TreeNode, root2 *TreeNode) bool {
    // 0 0
    if root1 == nil && root2 == nil {
        return true
    }
    // 0 1/1 0 |0 0 被排除
    if root1 == nil || root2 == nil {
        return false
    }
    if root1.Val != root2.Val {
        return false
    }
    return dfs(root1.Left, root2.Left) && dfs(root1.Right, root2.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val  {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
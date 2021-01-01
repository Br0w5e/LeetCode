package main

//235. 二叉搜索树的最近公共祖先

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	ancestor := root
	for {
		//在左子树
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
			//在右子树
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			//左右子树
			return ancestor
		}
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	if root == nil || p.Val <= root.Val && q.Val >= root.Val {
		return root
	} else if p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return lowestCommonAncestor(root.Left, p, q)
	}
}

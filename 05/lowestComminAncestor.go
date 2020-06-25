package main
//寻找二叉树的最近的共同祖先
//方法：
/*
如果p, q 分别位于左，右子树，则根即为要找的节点
如果p, q 都位于左子树，则公共祖先由左子树中产生。反之则从右子树中产生
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}else{
		return right
	}
}


//下边是一个二叉搜索树的最近祖先的C代码
/*
struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    if((root->val - p->val)*(root->val - q->val) <= 0 ){
        return root;
    }
    if(root->val > p->val){
        return lowestCommonAncestor(root->left, p, q);
    }
    if(root->val < p->val){
        return lowestCommonAncestor(root->right, p, q);
    }

    return NULL;
}
*/
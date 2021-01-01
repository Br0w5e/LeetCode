//判定给定的序列是不是搜索二叉树的一个遍历序列
//思路：当前队列的最后一个肯定是根节点
//迭代算法
func verifyPostorder(postorder[]int ) bool{
	n := len(postorder)
	if n  < 2 {
		return true
	}
	tail := n-1
	for tail != 0 {
		popinter := 0
		//左子树小于根节点
		for postorder[popinter] < postorder[tail] {
			popinter++
		}
		for popinter != tail {
			//右子树大于根节点
			if postorder[popinter] > postorder[tail]{
				popinter++
			} else {
				return false
			}
		}
		tail--
	}
	return true
}

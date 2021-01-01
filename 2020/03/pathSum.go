//返回二叉树中所有满足指定和的路径
func pathSum(root *TreeNode, sum int)  [][]int{
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root, sum, []int{}, &ret)
	return ret
}

func dfs(root *TreeNode, sum int, arr []int, ret *[][]int){
	if root == nil{
		return
	}
	arr = append(arr, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil{
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*ret = append(*ret, tmp)
	}
    dfs(root.Left,sum - root.Val,arr,ret)
	dfs(root.Right,sum - root.Val,arr,ret)
	arr = arr[:len(arr)-1]
}

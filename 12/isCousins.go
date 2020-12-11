package main

//993. 二叉树的堂兄弟节点


func isCousins(root *TreeNode, x int, y int) bool {
	dmap := map[int]int{}
	pmap := map[int]*TreeNode{}
	dfs(root,nil,0,&dmap,&pmap)

	return dmap[x] == dmap[y] && pmap[x] !=  pmap[y]
}
//计算depth
//创建parent map
func dfs(root,parent *TreeNode,d int,dm *map[int]int,pm *map[int]*TreeNode){
	if root == nil {
		return
	}
	(*dm)[root.Val]=d
	(*pm)[root.Val]=parent
	dfs(root.Left,root,d+1,dm,pm)
	dfs(root.Right,root,d+1,dm,pm)
}

//dfs 逐个遍历
func isCousins(root *TreeNode, x int, y int) bool {
	p1, d1 := 0, 0
	dfs(root, 0, 0, x, &p1, &d1)
	p2, d2 := 0, 0
	dfs(root, 0, 0, y, &p2, &d2)
	return p1 != p2 && d1 == d2
}

func dfs(root *TreeNode, parVal, depth, x int, p *int, d *int){
	if root.Val == x{
		*p, *d = parVal, depth
		return
	}
	if root.Left != nil {
		dfs(root.Left, root.Val, depth + 1, x, p, d)
	}
	if root.Right != nil {
		dfs(root.Right, root.Val, depth + 1, x, p, d)
	}
}
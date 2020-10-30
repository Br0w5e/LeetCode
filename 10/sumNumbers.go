package main

//129. 求根到叶子节点数字之和

//dfs
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

//前序遍历

func dfs(root *TreeNode, sum int) int {
	//递归出口
	if root == nil {
		return 0
	}
	//计算数值
	sum = sum*10+root.Val
	//返回当前路径的和
	if root.Left == nil && root.Right == nil {
		return sum
	}
	//递归计算和
	return dfs(root.Left, sum)+dfs(root.Right, sum)
}


//bfs
type pair struct {
	node *TreeNode
	num int
}

func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	queue := make([]pair, 0)
	queue = append(queue, pair{root, root.Val})
	for len(queue) > 0 {
		//出队
		p := queue[0]
		queue = queue[1:]
		left, right, num := p.node.Left, p.node.Right, p.num
		//判断是否为叶子节点
		if left == nil && right == nil {
			sum += num
		} else {
			//遍历，注意要加上本下一个节点的值
			if left != nil {
				queue = append(queue, pair{left, num*10+left.Val})
			}
			if right != nil {
				queue = append(queue, pair{right, num*10+right.Val})
			}
		}
	}
	return sum
}
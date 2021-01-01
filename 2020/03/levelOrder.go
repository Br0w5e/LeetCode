//按行打印树的各个节点
//思路：广度优先遍历，重点：设置额外队列进行处理
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := make([]int, 0)
	for len(queue) > 0{
		t := queue[0]
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
		res = append(res, t.Val)
		queue = queue[1:]
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)

    if root == nil{
        return res
    }
    queue := []*TreeNode{root}
    ans := 0
    for len(queue) > 0{
        n := len(queue)
        res = append(res, make([]int, 0))
        for i := 0; i < n; i++{
            top := queue[0]
            if top.Left != nil {
                queue = append(queue, top.Left)
            }
            if top.Right != nil {
                queue = append(queue, top.Right)
            }
            res[ans] = append(res[ans], top.Val)
            queue = queue[1:]
        }
        ans++
    }
    return res
}
//从下到上打印树
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil{
        return res
    }

    queue := []*TreeNode{root}
    ans := 0
    for len(queue) > 0{
        n := len(queue)
        res = append(res, make([]int, 0))
        for i := 0; i < n; i++{
            top := queue[0]
            if top.Left != nil {
                queue = append(queue, top.Left)
            }
            if top.Right != nil {
                queue = append(queue, top.Right)
            }
            res[ans] = append(res[ans], top.Val)
            queue = queue[1:]
        }
        ans++
    }
    
	left := 0
	right := len(res) - 1
	for left < right {
        res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return res
}
//之字形打印二叉树
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }

    queue := []*TreeNode{root}
    level := 0
    for len(queue) > 0 {
        tmp := []*TreeNode{}
        res = append(res, make([]int, 0))
        for _, v := range queue {
            if level % 2 == 0 {
                res[level] = append(res[level], v.Val)
            } else {
                s := []int{v.Val}
                res[level]=append(s,res[level]...)
            }
            if v.Left != nil {
                tmp = append(tmp, v.Left)
            }
            if v.Right != nil{
                tmp = append(tmp, v.Right)
            }
        }
        queue = tmp
        level++
    }
    return res
}
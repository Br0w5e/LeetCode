//计算盲人按摩师的最大工作时间
//动态规划

func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func massage(nums []int) int {
    a, b := 0, 0
    for i := 0; i < len(nums); i++{
        c := max(b, a+nums[i])
        a = b
        b = c
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//打家劫舍，一排
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = max(dp[0], nums[1])
    for i := 2; i < n; i++{
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }
    return dp[n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    preMax := 0 
    curMax := 0
    for _, v := range nums {
        temp := curMax 
        if preMax+v > curMax {
            curMax = preMax + v
        }
        preMax = temp
    }
    return curMax
}

//打家劫舍，成环
//思路：[:n-1],[1:n]取最大。即不抢第一个或者最后一个。
func rob(nums []int) int {
    n := len(nums)
    if n == 0{
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    n1 := rober(nums[:n-1])
    n2 := rober(nums[1:n])
    return max(n1, n2)
}

func rober(nums []int) int {
    preMax := 0 
    curMax := 0
    for _, v := range nums {
        temp := curMax 
        if preMax+v > curMax {
            curMax = preMax + v
        }
        preMax = temp
    }
    return curMax
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//打家劫舍：二叉树
//方法一：递归
func rob(root *TreeNode) int {
    if root == nil {
        return 0
    }
    money := root.Val
    if root.Left != nil {
        money += (rob(root.Left.Left) + rob(root.Left.Right))
    }
    if root.Right != nil {
        money +=(rob(root.Right.Left) + rob(root.Right.Right))
    }
    return max(money, rob(root.Left)+rob(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//方法二：树形DP。这个不太懂
func rob(root *TreeNode) int {
	mp := make(map[*TreeNode]int)
	return robImpl(root, mp)
}

func robImpl(root *TreeNode, mp map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if val, ok := mp[root]; ok {
		return val
	}
	money := root.Val
	if root.Left != nil {
		money += robImpl(root.Left.Left, mp) + robImpl(root.Left.Right, mp)
	}
	if root.Right != nil {
		money += robImpl(root.Right.Left, mp) + robImpl(root.Right.Right, mp)
	}
	r :=  max(money, robImpl(root.Left, mp) + robImpl(root.Right, mp))
	mp[root] = r
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
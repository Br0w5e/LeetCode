package main
//55. 跳跃游戏

//能否跳跃到最后位置
//从后往前遍历
func canJump(nums []int) bool {
    //到末尾的距离
    step := 1
    for i := len(nums)-2; i >= 0; i-- {
        //可以到达末尾，则目前的置位末尾
        if nums[i] >= step {
            step = 1
        } else {
        //不能到达末尾，则到达末尾的距离加1
            step++
        }
    }
    //最后到达末尾的距离是不是1
    return step == 1
}

//贪心
func canJump(nums []int) bool {
    n := len(nums)
    //最远可访问的距离
    last := 0
    for i := 0; i < n; i++{
        if i <= last {
			//更新最远距离
            last = max(last, i+nums[i])
            if last >= n-1 {
                return true
            }
        }
    }
    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
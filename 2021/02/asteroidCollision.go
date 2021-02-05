package main

//735. 行星碰撞

//分情况讨论，栈思想
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		//栈为空或者行星向右运行，入栈
		if len(stack) == 0 || asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
			//行星向左运行
		} else {
			//向左的体量大
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -asteroids[i] {
				stack = stack[:len(stack)-1]
			}
			//左右体量相等
			if len(stack) > 0 && stack[len(stack)-1] == -asteroids[i] {
				stack = stack[:len(stack)-1]
				continue
			}
			//栈为空，或者行栈顶星向左运行
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, asteroids[i])
			}
		}
	}
	return stack

}

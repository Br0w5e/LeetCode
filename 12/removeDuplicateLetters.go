package main

//316. 去除重复字母

//参考smallestSubsequence、1081. 不同字符的最小子序列
//单调栈+map
func removeDuplicateLetters(s string) string {
	// 每个字符出现次数
	count := make([]int, 26)
	// 是否在栈中，存在为true
	exist := make([]bool, 26)
	// 单调栈
	stack := make([]byte, 0)

	// 统计字符出现次数
	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range s {
		// 栈中已有c，跳过
		if exist[c-'a'] {
			// 同时减少这个字符出现的次数
			count[c-'a']--
			continue
		}

		// 出栈的核心判断要素：
		// 1. 栈里有元素
		// 2. 栈顶元素大于当前元素c
		// 3. 栈顶元素在后续出现
		for len(stack) > 0 && stack[len(stack)-1] > byte(c) && count[stack[len(stack)-1]-'a'] > 0 {
			// 进入这里说明栈顶元素大于当前元素，所以不符合字典序，要把栈顶元素出栈
			// 【重要】for 循环确保栈里面保存的都是比当前元素小的，因为大的都被pop掉了

			// 标记为栈中不含栈顶元素
			exist[stack[len(stack)-1]-'a'] = false
			// 删除栈顶元素
			stack = stack[:len(stack)-1]
		}

		// 添加新字符
		stack = append(stack, byte(c))
		// 减少该字符出现次数
		count[c-'a']--
		// 标记栈中有该字符
		exist[c-'a'] = true
	}

	return string(stack)
}

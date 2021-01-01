package main

//842. 将数组拆分成斐波那契序列

import "math"

func splitIntoFibonacci(s string) []int {
	n := len(s)
	F := make([]int, 0)
	var backTrack func(index, sum, prev int) bool
	backTrack = func(index, sum, prev int) bool {
		// 如果遍历到最后,F中的长度大于等于3，则满足条件
		if index == n {
			return len(F) >= 3
		}
		// 当前判断的数字
		cur := 0

		// 从当前索引位置开始，判断后面的数能不能组成斐波那契数列
		for i := index; i < n; i++ {
			// 除第一个数字外，其余数字不能以0开头
			if i > index && s[index] == '0' {
				break
			}

			cur = cur*10 + int(s[i]-'0')
			// 拆出的整数要符合 32 位有符号整数类型
			if cur > math.MaxInt32 {
				break
			}

			// 如果F的长度小于2，表示该斐波那契数列前两项还未添加，需要先添加前两项
			if len(F) >= 2 {
				// 如果当前值不等于前两个数之和
				// 如果当前数小于前两个数之和，还要继续判断取下一位增加当前数，进行判断，直接进行下一次循环
				if cur < sum {
					continue
				}
				// 如果当前数大于前两个数之和，就不用继续判断了，无法构成斐波那契数列直接退出
				if cur > sum {
					break
				}
				// 如果不大于，也不小于则刚好满足条件
			}
			// 要么斐波那契数列前两项未添加完成，要么cur等于前两项之和，cur复合要求，加入序列 F
			F = append(F, cur)
			// 进行递归对后面的数据进行判断
			if backTrack(i+1, prev+cur, cur) {
				return true
			}
			// 前一项不满足条件，回退还原状态，进行下一次循环，更改前一项
			F = F[:len(F)-1]
		}
		return false
	}
	backTrack(0, 0, 0)
	return F
}

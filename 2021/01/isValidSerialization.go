package main

import "strings"

//331. 验证二叉树的前序序列化

// s从后遍历，用num记录#的个数，当遇到正常节点时，#的个数-2，并将该节点转化成#，num+1，，整体即为num-1; 当出现#的个数不足2时，即false。最终只有根节点一个#
func isValidSerialization(s string) bool {
	num := 0
	for i := len(s)-1; i >= 0; i-- {
		//扫描遇到逗号
		if s[i] == ',' {
			continue
		}
		//遇见#，num加1
		if s[i] == '#' {
			num++
		} else {
			//遇见数字，需要进行判断，能不能继续合并
			for i >= 0 && s[i] != ',' {
				i--
			}
			//能合并，减去两个#， 并将本节点置为#，相当于减去1个#
			if num >= 2 {
				num--
			} else {
				return false
			}
		}
	}
	//最终只有根节点一个#
	return num == 1
}

//数组化表示
func isValidSerialization2(preorder string) bool {
	arr := strings.Split(preorder, ",")
	num := 0
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] == "#" {
			num++
		} else {
			if num >= 2 {
				num--
			} else {
				return false
			}
		}
	}
	return num == 1
}
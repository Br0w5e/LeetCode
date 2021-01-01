# 04.09

[TOC]

## [面试题 03.04. 化栈为队](https://leetcode-cn.com/problems/implement-queue-using-stacks-lcci/)

**思路：**
- 用栈实现队列
- 一个栈就行了

**代码：**

- [点击这里](./MyQueue.go)

## [657. 机器人能否返回原点](https://leetcode-cn.com/problems/robot-return-to-origin/)

**思路：**
- 向左走的步数和向右走的步数应该是一样的，向下走的步数和和向上走的步数一样
- 用单个flag或者多个flag进行标记操作

**代码：**
- [点击这里](./judgeCircle.go)

## [1309. 解码字母到整数映射](https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/)

**思路：**
- 先考虑3个的，再考虑一个的
- 注意映射关系
**代码：**
- [点击这里](./freqAlphabets.go)

## [1370. 上升下降字符串](https://leetcode-cn.com/problems/increasing-decreasing-string/)

**思路：**
- 统计每个字母出现的次数；然后升序走一遍，降序走一遍，知道统计值为0

**代码：**
- [点击这里](./sortString.go)

## [1304. 和为零的N个唯一整数](https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/)

**思路：**
- 方法很多的
- 对称实现
- 最后一个反向求和实现

**代码：**
- [点击这里](./sumZero.go)

## [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

**思路：**
- dfs
- 注意剪枝
- 注意递归出口

**代码：**
- [点击这里](./generateParenthesis.go)

## [面试题 01.02. 判定是否互为字符重排](https://leetcode-cn.com/problems/check-permutation-lcci/)

**思路：**

- 利用一个共享数组或者map，在s1中出现的字符+1，在s2中出现的字符-1
- 判断最后数组是不是全0

**代码：**

- [点击这里](./CheckPermutation.go)


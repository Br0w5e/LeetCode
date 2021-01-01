# 04.06

[TOC]

## [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)

**思路：**

- 经典动态规划
- 不会看一下题解，没什么讲的。重点搞一下状态转移方程，代码背诵。
- 若当前字母相同，则`dp[i][j] = dp[i - 1][j - 1]`;
- 否则取增\删、替、三个操作的最小值 + 1， 即:    `dp[i][j] = min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]) + 1`

**代码：**

- [点击这里](./minDistance.go)

## [面试题 16.01. 交换数字](https://leetcode-cn.com/problems/swap-numbers-lcci/)

**思路：**

- 不使用中间变量，go中有语法糖的；
- 加法实现
- 异或实现

**代码：**

- [点击这里](./swapNumbers.go)

## [面试题 08.07. 无重复字符串的排列组合](https://leetcode-cn.com/problems/permutation-i-lcci/)

**思路：**

- DFS
- 回溯

**代码：**

- [点击这里](./permutation.go)
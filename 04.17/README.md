# 04.17

[TOC]

## [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

**思路：**

- 从后往前考虑，如果数组上的数字大于到结尾的步数，则将当前位置标记为结尾，否则向前遍历并将步伐加一
- 最终判断第一个位置上的是不是距离最终结尾等于1
- 方法二：贪心算法：每次标记最远可访问的位置，判断最终该位置是不是大于最后位置。

**代码：**

- [点击这里](./canJump.go)

## [748. 最短完整词](https://leetcode-cn.com/problems/shortest-completing-word/)

**思路：**

- 将牌照统一转化为小写；
- 写一个match函数，map牌照中的字母；
- 用单词去匹配map，判断其中的value值是不是等于0；
- 比较现有的和以前的长度取最小的。

**代码：**

- [点击这里](./shortestCompletingWord.go)

## [1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)

**思路：**

- 二维dp，参考一下编辑距离；
- dp关系式：`A[i] == B[j] --- > dp[i+1][j+1] = dp[i][j]+1` 否则`dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])` 

**代码：**

- [点击这里](./maxUncrossedLines.go)

## [1035. 不相交的线](https://leetcode-cn.com/problems/uncrossed-lines/)

**思路：**

- 相同的成对，并求最多的不想交线，翻译出来也是最长子序列问题；
- 参考最长子序列，dp关系式和最长子序列一样；

**代码：**

- [点击这里](./maxUncrossedLines.go)

## [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

**思路：**

- 相当于求最大上升差，不算下降的即可

**代码：**

- [点击这里](./maxProfit.go)


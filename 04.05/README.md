# 04.05

[TOC]

记录双周赛和单周赛的一些题

## [5361. 圆和矩形是否有重叠](https://leetcode-cn.com/problems/circle-and-rectangle-overlapping/)

**思路：**

- 判断矩形边上各点和圆心距离与半径的大小关系；
- 书写方便，直接判断所有整数点。

**代码：**

- [点击这里](./checkOverlap.go)

## [5360. 统计最大组的数目](https://leetcode-cn.com/problems/count-largest-group/)

**思路：**

这题建议去看一下英文解释，反正汉语我不好懂

- 模拟就完事了。

**代码：**

- [点击这里](./countLargestGroup.go)

## [5362. 构造 K 个回文字符串](https://leetcode-cn.com/problems/construct-k-palindrome-strings/)

**思路：**

出现奇数次字母的个数决定能组成多少个回文串。（证明网上找一下）

- s可以构成k个非空回文串当且仅当c <= k <= len(s)，其中c为出现奇数次不同字母的个数。
- 因为出现偶数次的字母可以组成len(s)个回文串

**代码：**

- [点击这里](./canConstruct.go)

## [5376. 非递增顺序的最小子序列](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order/)

**思路：**

- 先排序
- 再暴力走一趟

**代码：**

- [点击这里](./minSubsequence.go)

## [5377. 将二进制表示减到 1 的步骤数](https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/)

**思路：**

- 先转化为10进制然后暴力模拟——>超时
- 直接操作字符串——>t通过

**代码：**

- [点击这里](./numSteps.go)
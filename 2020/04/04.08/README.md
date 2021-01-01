# 04.08

[TOC]

## [面试题13. 机器人的运动范围](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

**思路：**

- DFS
- 加判断条件和辅助标记数组

**代码：**

- [点击这里](./movingCount.go)

## [709. 转换成小写字母](https://leetcode-cn.com/problems/to-lower-case)

**思路：**

- 逐个遍历遇见大写+32
- 位运算

- 大写变小写、小写变大写 : 字符 ^= 32;
- 大写变小写、小写变小写 : 字符 |= 32;
- 小写变大写、大写变大写 : 字符 &= -33;

**代码：**

- [点击这里](./toLowerCase.go)

## [1323. 6 和 9 组成的最大数字](https://leetcode-cn.com/problems/maximum-69-number/)

**思路：**

- 先将数字转换为字符串或者数组，然后将做高位的6变成9，再转化回来。

**代码：**

- [点击这里](./maximum69Number.go)

## [461. 汉明距离](https://leetcode-cn.com/problems/hamming-distance/)

**思路：**

- 将`x`和`y`进行异或操作，得到的数据转化为二进制，其中1的个数就是不同的位数

**代码：**

- [点击这里](./hammingDistance.go)

## [617. 合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/)

**思路：**

- 递归
- 确定好递归出口

**代码：**

- [点击这里](./mergeTrees.go)

## [938. 二叉搜索树的范围和](https://leetcode-cn.com/problems/range-sum-of-bst/)

**思路：**

- dfs

- 当前搜索值小于L时候搜索其右子树，当前搜索值大于于R时候搜索其左子树。

**代码：**

- [点击这里](./rangeSumBST.go)
# 04.03
[TOC]
## [寻找缺失的数字](https://leetcode-cn.com/problems/missing-number-lcci/)
**思路：**

题目要求：是求N（包括N）以内中缺失的数字，我们可以采用等差数列求和得出N包括N的和然后求数组中所有数字的和。两者之差就是最终结果。

借助range性质，和0的性质我们就可以在一次遍历情况下求得最终的结果

**代码：**

- [点击这里](./missingNumber.go)

## [倒数第K个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)
**思路：**

- 遍历所有链表节点，然后进行最后一个判断
- 设置快慢指针，让快指针先走K步，然后快慢指针同时走，当快指针走向终点的时候，慢指针走向倒数第K个节点。 

**代码：**
- [点击这里](./kthToLast.go)

## [字符串转数字](https://leetcode-cn.com/problems/string-to-integer-atoi/)
**思路：**
- 遍历，遇见不同的情况进行不同的处理
- 注意整数越界时候的处理

**代码：**
- [点击这里](./myAtoi.go)


## [最小差值](https://leetcode-cn.com/problems/smallest-range-i/)

**思路：**
题目解读：给数组中添加[-K, K]中的任意数组保证数组里边所以数差值最小，并求该差值

**代码：**
- [点击这里](./smallestRangeI.go)
# 04.04

[TOC]

## [1021. 删除最外层的括号](https://leetcode-cn.com/problems/remove-outermost-parentheses)

**思路：**

- 利用一个flag位，统计括号，最后将外层括号去掉加入结果中

**代码：**

- [点击这里](./removeOuterParentheses.go)

## [1221. 分割平衡字符串](https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/)

**思路：**

- 利用flag记录`L`或者`R`
- 当flag == 0的时候记录值加一

**代码：**

- [点击这里](./balancedStringSplit.go)

## [1299. 将每个元素替换为右侧最大元素](https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/)

**思路：**

- 从右往左遍历，tmp记录前一状态最大值，并求当前最大值，然后进行置换

**代码：**

- [点击这里](./replaceElements.go)

## [1351. 统计有序矩阵中的负数](https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/)

**思路：**

- 从又上或者左下起开始进行统计

- 偷懒直接遍历

**代码：**

- [点击这里](./countNegatives.go)

## [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

**思路：**
- 将柱子填充成左边不严格递增，右边不严格递增的图形，最终增量即为接到的水的数量
- 双指针

**代码：**
- [点击这里](./trap.go)
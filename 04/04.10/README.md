# 04.10

[TOC]

## [面试题 08.04. 幂集](https://leetcode-cn.com/problems/power-set-lcci/)、[78. 子集](https://leetcode-cn.com/problems/subsets/)

**思路：**

- 这两个题是同一个
- 遍历给定数组，遇到一个数就把当前所有子集加上该数组成新的子集，遍历完毕即是所有子集。

**代码：**

- [点击这里](./subsets.go)

## [338. 比特位计数](https://leetcode-cn.com/problems/counting-bits/)

**思路：**

- 逐个遍历，遇到每个数计算其二进制中1的个数
- 利用i和i/2中1的个数进行处理：
- i为偶数i和i/2中1的个数相同，i为奇数i比i/2多一个1

**代码：**

- [点击这里](./countBits.go)

## [面试题56 - II. 数组中数字出现的次数 II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)

**思路：**

- 生成map去判断，寻找值为1的键
- 利用异或运算

**代码：**

- [点击这里](./singleNumber.go)
- [出现一次的数](./onlyone.go)

## [807. 保持城市天际线](https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/)

**思路：**

- 将新矩阵置位行最大矩阵

- 根据列最大值修正新矩阵

- 求新矩阵中每个元素和原来元素的差值，并求和返回

**代码：**

- [点击这里](./maxIncreaseKeepingSkyline.go)

## [151. 翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

**思路：**

- 库函数使用`Split`和`Join`

**代码：**

- [点击这里](./reverseWords.go)
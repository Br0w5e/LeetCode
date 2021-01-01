# 04.15

[TOC]

## [1337. 方阵中战斗力最弱的 K 行](https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/)

**思路：**

- 计算1的个数
- 用1的个数乘以最大的军人数，加上行号
- 对数组排序，并对前k个对上一步的乘数取余

**代码：**

- [点击这里](./kWeakestRows.go)

## [766. 托普利茨矩阵](https://leetcode-cn.com/problems/toeplitz-matrix/)

**思路：**

- 弄清对角线上的位置关系
- 遍历

**代码：**

- [点击这里](./isToeplitzMatrix.go)

## [1200. 最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference/)

**思路：**

- 求最小绝对值差
- 求解数组

**代码：**

- [点击这里](./minimumAbsDifference.go)

## [面试题 08.10. 颜色填充](https://leetcode-cn.com/problems/color-fill-lcci/)

**思路：**

- dfs
- 方向控制，好像以前做过

**代码：**

- [点击这里](./floodFill.go)

## [542. 01 矩阵](https://leetcode-cn.com/problems/01-matrix/)

**思路：**

- 设起始 level = 0

- 遍历矩阵遇到0就把4周的1设置为 level - 1

- level--

- 继续第二次遍历 直到没有新level

**代码：**

- [点击这里](./updateMatrix.go)


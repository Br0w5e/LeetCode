# 04.02
本次主要练习方向控制类。
[TOC]
## [289.生命游戏](https://leetcode-cn.com/problems/game-of-life/)
**思路：**

- 设置方向向量
- 可以重新开辟空间进行做
- 注意每个细胞状态同时变化，因此不能一个一个更新
- 应用最低位存储当前状态，倒数第二位存储下一状态

**代码**：

- [点击这里](./gameOfLife.go)

## [59.螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

**思路：**

- 带有方向控制，逐个生成

**代码：**

- [点击这里](./generateMatrix.go)

## [200.岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

**思路：**

- DFS
- 结合方向控制

**代码：**
- [点击这里](./numIslands.go)

## [892.三维形体的表面积](https://leetcode-cn.com/problems/surface-area-of-3d-shapes/)

**思路：**

- 每次总是多两个上下底面
- 侧面多少由相邻四个高度决定

**代码：**

- [点击这里](./surfaceArea.go)

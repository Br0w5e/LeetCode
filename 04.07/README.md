# 04.07

[TOC]

## [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array)
**思路：**
- 带优化的暴力模拟
- 三次翻转，前边k部分、后边k部分、整体

**代码：**
- [点击这里](./rotate1.go)
## [733. 图像渲染](https://leetcode-cn.com/problems/flood-fill)
**思路：**
- dfs
- 终止条件，颜色跟给定颜色不同
- 方向控制

**代码：**
- [点击这里](./floodFill.go)
## [835. 图像重叠](https://leetcode-cn.com/problems/image-overlap)
**思路：**
- 暴力走一趟，将重叠的一计数并返回
- 不断调整最大的

**代码：**
- [点击这里](./largestOverlap.go)
## [832. 翻转图像](https://leetcode-cn.com/problems/flipping-an-image)
**思路：**
- 暴力模拟
- 翻转的时候就进行异或操作

**代码：**
- [点击这里](./flipAndInvertImage.go)


## [面试题 01.07. 旋转矩阵](https://leetcode-cn.com/problems/rotate-matrix-lcci)
**思路：**
- 寻找旋转后的坐标对应关系 `matrix[i][j]=matrix[j][n-1-i]`
- 进行思路总结：先翻转，在对折
- 或者原地旋转

**代码：**
- [点击这里](./rotate.go)
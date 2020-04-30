# 04.16

[TOC]

## [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)

**思路：**

- 将区间按照左端点排序
- 逐步向后合并，不能合并就将其加入结果中
- 知道最后一个

**代码：**

- [点击这里](./merge.go)

## [63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

**思路：**

- DP，可在原数组上修改
- 特殊处理第一行和第一列
- 其他`dp[i][j] = dp[i-1][j]+dp[i][j-1]`

**代码：**

- [点击这里](./uniquePathsWithObstacles.go)

## [524. 通过删除字母匹配到字典里最长单词](https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/)

**思路：**

- 先写一个包含函数，使用双指针
- 逐个比较数组中的字符串，注意字典顺序

**代码：**

- [点击这里](./findLongestWord.go)

## [575. 分糖果](https://leetcode-cn.com/problems/distribute-candies/)

**思路：**

- 计算糖果的种类
- 返回种类和数目一半的最小值

**代码：**

- [点击这里](./distributeCandies.go)

## [852. 山脉数组的峰顶索引](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/)

**思路：**

- 线性扫描
- 二分扫描（注意条件）

**代码：**

- [点击这里](./peakIndexInMountainArray.go)
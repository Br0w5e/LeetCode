package main

//暴力解法
func countSmaller(nums []int) []int {
	res := make([]int, 0)
	for i := 0 ; i < len(nums); i++ {
		count := 0
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

//树状数组
type TreeNode struct {
	Val int
	Count int
	Left *TreeNode
	Right *TreeNode
}
func countSmaller1(nums []int) []int {
	length := len(nums)
	count := make([]int, length)
	if length <= 1 {
		return count
	}
	
	//从右往左处理
	node := TreeNode{Val: nums[length-1]}
	count[length-1] = 0

	for i := length-2; i >= 0 ; i-- {
		var data int
		BST_insert(&node, &TreeNode{Val: nums[i]}, &data)
		count[i] = data
	}
	return count
}

func BST_insert(node *TreeNode, insert_node *TreeNode, small_count *int) {
	//*放在左边
	if insert_node.Val <= node.Val {
		node.Count++
		if node.Left != nil {
			BST_insert(node.Left, insert_node, small_count)
		} else {
			node.Left = insert_node
		}
	}
	if insert_node.Val > node.Val {
		*small_count = *small_count+node.Count+1
		if node.Right != nil {
			BST_insert(node.Right, insert_node, small_count)
		} else {
			node.Right = insert_node
		}
	}
}

//归并排序
var(
	//临时数组
	tmp   []int
	//存放虚拟索引
	index []int
	//存放结果值
	count []int
)

//归并排序实现
func countSmaller2(nums []int) []int {
	if nums==nil || len(nums)==0{
		return nil
	}
	tmp=make([]int,len(nums))
	count=make([]int,len(nums))
	//存放虚拟索引
	index=make([]int,len(nums))
	//不改变nums的值，通过虚拟索引进行归并排序
	for i:=0;i<len(nums);i++{
		index[i]=i//0，1，2，3分别代表5，2，6，1
	}
	mergesort(nums,0,len(nums)-1)

	return count
}

func mergesort(nums []int,left,right int){
	if left>=right{
		return
	}
	mid:=(right+left)/2
	mergesort(nums,left,mid)
	mergesort(nums,mid+1,right)
	merge(nums,left,mid,right)
}

func merge(nums []int,left,mid,right int){
	//index相当于nums数组
	l:=left
	r:=mid+1
	k:=left
	for l<=mid && r<=right{
		//当左边出列时，此时的逆序对的长度就是右边出去的个数
		if nums[index[l]] <= nums[index[r]]{
			tmp[k]=index[l]
			//长度为r-1-mid:当右边出列后，r++,指向下一个，所以实际为
			//r-1,mid和r-1之间的数量为r-1-mid
			//得到结果后加入到count中，此时加入到count中当前出列元素所对应的索引下
			count[index[l]]+=(r-1-mid)
			l++
			k++
		}else{//右边小则出去
			tmp[k]=index[r]
			r++
			k++
		}
	}
	for l<=mid{
		tmp[k]=index[l]
		count[index[l]]+=(r-1-mid)
		k++
		l++
	}
	for r<=right{
		tmp[k]=index[r]
		r++
		k++
	}
	for i:=left;i<=right;i++{
		index[i]=tmp[i]
	}
}
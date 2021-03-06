package main

//303. 区域和检索 - 数组不可变

type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	res := new(NumArray)
	res.nums = nums
	return *res
}


func (this *NumArray) SumRange(i int, j int) int {
	res := 0
	for _, v := range this.nums[i:j+1] {
		res += v
	}
	return res
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */


type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for i, num := range nums {
		arr[i+1] = arr[i]+num
	}
	return NumArray{
		nums: arr,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j+1]-this.nums[i]
}
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
